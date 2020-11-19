package db

// Health checks main connectivity to Cassandra
func (d *DB) Health() (bool, error) {
	sess, err := d.Session()
	if err != nil {
		return false, err
	}

	q := sess.Query("SELECT now() FROM system.local")
	if err = q.Exec(); err != nil {
		return false, err
	}

	return true, nil
}
