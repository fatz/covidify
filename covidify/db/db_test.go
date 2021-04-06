package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"testing"

	mysqltest "github.com/lestrrat-go/test-mysqld"
	"github.com/stretchr/testify/assert"
)

var mysqld *mysqltest.TestMysqld
var db *sql.DB
var dsn string

func setup(dbname string) (*sql.DB, string, error) {
	var err error

	config := mysqltest.NewConfig()
	config.SkipNetworking = false
	config.Port = 13306

	// Starts mysqld listening on port 13306

	mysqld, err = mysqltest.NewMysqld(config)
	if err != nil {
		return nil, "", err
	}
	db, err = sql.Open("mysql", mysqld.DSN())
	if err != nil {
		return nil, "", err
	}
	_, err = db.Query(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbname))
	if err != nil {
		return nil, "", err
	}

	_, err = db.Query(fmt.Sprintf("CREATE DATABASE `%s`;", dbname))
	if err != nil {
		return nil, "", err
	}

	dsn = mysqld.DSN(mysqltest.WithDbname(dbname))
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, "", err
	}

	schema, err := ioutil.ReadFile(fmt.Sprintf("../../%s.sql", dbname))
	if err != nil {
		return nil, "", err
	}
	schemaStatements := strings.Split(string(schema), ";")

	for _, schemaStatement := range schemaStatements {
		if strings.TrimSpace(schemaStatement) == "" {
			continue
		}
		_, err = db.Exec(schemaStatement)
		if err != nil {
			return nil, "", err
		}
	}

	// testdata, err := ioutil.ReadFile("./testdata.sql")
	//
	// if err != nil {
	// 	return nil, "", err
	// }
	//
	// testdataStatements := strings.Split(string(testdata), ";")
	//
	// for _, testdataStatement := range testdataStatements {
	// 	if strings.TrimSpace(testdataStatement) == "" {
	// 		continue
	// 	}
	// 	_, err = db.Query(testdataStatement)
	// 	if err != nil {
	// 		return nil, "", err
	// 	}
	// }

	return db, dsn, err
}

func TestMain(m *testing.M) {
	var err error

	db, dsn, err := setup("covidify")
	if err != nil {
		fmt.Printf("Setup Error (%v, %s) %v", db, dsn, err)
		os.Exit(1)
	}

	code := m.Run()
	mysqld.Stop()
	os.Exit(code)
}

func TestNewDB(t *testing.T) {
	d, err := NewDB(dsn)

	url, err := url.Parse(d.dsn)
	assert.NoError(t, err)

	assert.Equal(t, "true", url.Query().Get("parseTime"))

}
