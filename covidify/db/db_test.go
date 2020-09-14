package db

import (
	"fmt"
	"os"
	"strings"
)

func envConnect() (*DB, error) {
	c := os.Getenv("COVIDIFY_TEST_CLUSTER")

	if c == "" {
		return nil, fmt.Errorf("COVIDIFY_TEST_CLUSTER empty")
	}

	var k string
	if k = os.Getenv("COVIDIFY_TEST_KEYSPACE"); k == "" {
		k = "covidify"
	}

	return NewDB(strings.Split(c, ","), k)
}
