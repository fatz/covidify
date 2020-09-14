package covidify

import (
	"fmt"
	"strings"
)

type Config struct {
	CassandraConnection string
	CassandraKeyspace   string
	Port                *int
	Bind                string
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

	return c
}

func (c *Config) GetCassandraCluster() []string {
	return strings.Split(c.CassandraConnection, ",")
}

func (c *Config) GetBind() string {

	return fmt.Sprintf("%s:%d", c.Bind, c.Port)
}
