package covidify

import (
	"fmt"
	"net/http"
	"time"

	cdb "github.com/fatz/covidify/covidify/db"
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"github.com/etsy/statsd/examples/go"
)

type Server struct {
	config *Config
	db     *cdb.DB

	g *gin.Engine

	statsd *statsd.StatsdClient
}

func NewServerWithConfig(c *Config) (s *Server, err error) {
	s = new(Server)
	gin.SetMode(gin.ReleaseMode)

	s.config = c

	switch {
	case c.CassandraUsername != "" && c.CassandraPassword != "":
		s.db, err = cdb.NewDBWithPW(c.GetCassandraCluster(), c.CassandraKeyspace, c.CassandraUsername, c.CassandraPassword)
	default:
		s.db, err = cdb.NewDB(c.GetCassandraCluster(), c.CassandraKeyspace)
	}
	if err != nil {
		return nil, err
	}

	s.statsd = statsd.New(c.StatsDHost, c.StatsDPort)
	s.g = s.NewRouter()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(s.g)

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

	return svr.ListenAndServe()
}

// Health - Server health status
func (s *Server) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
