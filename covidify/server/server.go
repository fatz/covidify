package covidify

import (
	"net/http"
	"time"

	cdb "github.com/fatz/covidify/covidify/db"
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
)

type Server struct {
	config *Config
	db     *cdb.DB

	g *gin.Engine
}

func NewServerWithConfig(c *Config) (s *Server, err error) {
	s = new(Server)

	s.config = c

	s.db, err = cdb.NewDB(c.GetCassandraCluster(), c.CassandraKeyspace)
	if err != nil {
		return nil, err
	}

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

	svr := &http.Server{
		Addr:           ":8080",
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
