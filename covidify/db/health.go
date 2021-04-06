package db

// Health checks main connectivity to Cassandra
func (d *DB) Health() (bool, error) {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return false, err
	}

	if err := sqlDB.Ping(); err != nil {
		return false, err
	}

	return true, nil
}
