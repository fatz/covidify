package db

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

type DB struct {
	Cluster        []string
	Authencitcator gocql.Authenticator
	Keyspace       string

	ClusterConfig *gocql.ClusterConfig
	CqlSession    *gocql.Session
}

// NewDB does a simply connect to Cassandra by cluster address and keyspace
func NewDB(cluster []string, keyspace string) (db *DB, err error) {
	db = new(DB)

	if len(cluster) < 1 {
		return nil, fmt.Errorf("Specify at least one cluster node")
	}

	db.Cluster = cluster
	db.Keyspace = keyspace
	db.Authencitcator = nil

	return db.Connect()
}

// NewDBWithPW does a simply connect to Cassandra by cluster address and keyspace using username and password
func NewDBWithPW(cluster []string, keyspace, username, password string) (db *DB, err error) {
	db = new(DB)

	if len(cluster) < 1 {
		return nil, fmt.Errorf("Specify at least one cluster node")
	}

	db.Cluster = cluster
	db.Keyspace = keyspace

	db.Authencitcator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}

	return db.Connect()
}

func (d *DB) Connect() (*DB, error) {
	d.ClusterConfig = gocql.NewCluster(d.Cluster...)
	d.ClusterConfig.Keyspace = d.Keyspace
	d.ClusterConfig.Consistency = gocql.Quorum
	d.ClusterConfig.Timeout = 10 * time.Second

	if d.Authencitcator != nil {
		d.ClusterConfig.Authenticator = d.Authencitcator
	}

	_, err := d.Session()

	return d, err
}

func (d *DB) Session() (sess *gocql.Session, err error) {
	if d.CqlSession == nil {
		d.CqlSession, err = d.ClusterConfig.CreateSession()
		if err != nil {
			d.CqlSession = nil
			return nil, err
		}
	}

	// Check session state

	return d.CqlSession, nil
}
