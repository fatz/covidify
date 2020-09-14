package db

import (
	"fmt"

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

	return db.Connect()
}

func (d *DB) Connect() (*DB, error) {
	d.ClusterConfig = gocql.NewCluster(d.Cluster...)
	d.ClusterConfig.Keyspace = d.Keyspace
	d.ClusterConfig.Consistency = gocql.Quorum

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
