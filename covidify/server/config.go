package covidify

import (
	"fmt"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	MySQLHost     string
	MySQLPort     int
	MySQLUser     string
	MySQLPassword string
	MySQLDatabase string
	Port          *int
	Bind          string
	StatsDHost    string
	StatsDPort    int
	StatsDPrefix  string

	PrometheusPort        int
	PrometheusStandalone  bool
	PrometheusMetricsPath string

	Logger *log.Logger
}

func NewConfig() *Config {
	c := new(Config)

	//sample
	c.MySQLHost = "127.0.0.1"
	c.MySQLPort = 3306
	c.MySQLDatabase = "covidify"
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

func (c *Config) GetBind() string {

	return fmt.Sprintf("%s:%d", c.Bind, c.Port)
}

func (c *Config) GenDBDSN() string {
	u := url.URL{}
	u.Scheme = ""
	u.Host = fmt.Sprintf("tcp(%s:%d)", c.MySQLHost, c.MySQLPort)

	if c.MySQLUser != "" {
		u.User = url.UserPassword(c.MySQLUser, c.MySQLPassword)
	}
	u.Path = c.MySQLDatabase

	//skip scheme
	return u.String()[2:]
}
