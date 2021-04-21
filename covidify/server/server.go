package covidify

import (
	"fmt"
	"net/http"
	"time"

	statsd "github.com/etsy/statsd/examples/go"
	cdb "github.com/fatz/covidify/covidify/db"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type Server struct {
	config *Config
	db     *cdb.DB

	g *gin.Engine
	// Prometheus
	p *gin.Engine

	statsd *statsd.StatsdClient
}

func NewServerWithConfig(c *Config) (s *Server, err error) {
	s = new(Server)
	gin.SetMode(gin.ReleaseMode)

	s.config = c

	s.db, err = cdb.NewDB(s.config.GenDBDSN())
	if err != nil {
		return nil, err
	}

	s.statsd = statsd.New(c.StatsDHost, c.StatsDPort)
	p := ginprometheus.NewPrometheus("gin")
	p.MetricsPath = s.config.PrometheusMetricsPath

	s.g = s.NewRouterWithMiddleware([]gin.HandlerFunc{p.HandlerFunc()}, nil, nil)

	if s.config.PrometheusStandalone {
		// Start New Engine for standalone
		s.p = s.NewPrometheusRouter()
		p.SetMetricsPath(s.p)
	} else {
		// use default router
		p.SetMetricsPath(s.g)
	}

	return s, nil
}

func NewServer() (s *Server, err error) {
	c := NewConfig()

	return NewServerWithConfig(c)
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", s.config.Bind, *s.config.Port)
	svr := &http.Server{
		Addr:           addr,
		Handler:        s.g,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if s.p != nil {
		promAddr := fmt.Sprintf("%s:%d", s.config.Bind, s.config.PrometheusPort)
		promSvr := &http.Server{
			Addr:           promAddr,
			Handler:        s.p,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		// start standalone Prometheus router
		s.config.Logger.Infof("Starting Prometheus server: %s", promAddr)
		go func() { promSvr.ListenAndServe() }()
	}

	s.config.Logger.Infof("Starting server: %s", addr)
	return svr.ListenAndServe()
}

type Health map[string]bool

func (h *Health) ReturnCode() int {
	for _, s := range *h {
		if !s {
			return http.StatusInternalServerError
		}
	}
	return http.StatusOK
}

// Health - Server health status
func (s *Server) Health(c *gin.Context) {
	h := make(Health)

	dbOK, err := s.db.Health()
	if err != nil {
		s.config.Logger.Errorf("DB Healthy: %t - %s", dbOK, err.Error())
	}
	h["DB"] = dbOK

	c.JSON(h.ReturnCode(), h)
}
