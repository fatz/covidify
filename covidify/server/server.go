package covidify

import (
	"net/http"

	cdb "github.com/fatz/covidify/covidify/db"
	"github.com/gin-gonic/gin"
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

	return s, nil
}

func NewServer() (s *Server, err error) {
	c := NewConfig()

	return NewServerWithConfig(c)
}

func (s *Server) Run() error {

	return s.g.Run()
}

// Health - Server health status
func (s *Server) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
