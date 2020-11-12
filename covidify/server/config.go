package covidify

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	CassandraConnection string
	CassandraKeyspace   string
	CassandraUsername   string
	CassandraPassword   string
	Port                *int
	Bind                string
	StatsDHost          string
	StatsDPort          int
	StatsDPrefix        string

	PrometheusPort        int
	PrometheusStandalone  bool
	PrometheusMetricsPath string

	Logger *log.Logger
}

func NewConfig() *Config {
	c := new(Config)

	//sample
	c.CassandraConnection = "127.0.0.1"
	c.CassandraKeyspace = "covidify"
	var p int
	p = 3000
	c.Port = &p
	c.Bind = ""

	c.StatsDHost = "localhost"
	c.StatsDPort = 8125

	c.PrometheusPort = 8081
	c.PrometheusStandalone = true
	c.PrometheusMetricsPath = "/metrics"

	c.Logger = log.New()
	return c
}

func (c *Config) GetCassandraCluster() []string {
	return strings.Split(c.CassandraConnection, ",")
}

func (c *Config) GetBind() string {

	return fmt.Sprintf("%s:%d", c.Bind, c.Port)
}
