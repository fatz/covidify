package db

import (
	"log"
	"net/url"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	dsn string
	DB  *gorm.DB
}

// NewDB does a simply connect to Cassandra by cluster address and keyspace
func NewDB(dsn string) (db *DB, err error) {
	db = new(DB)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Warn, // Log level
			Colorful:      false,       // Disable color
		},
	)

	parsedDSN, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	if query := parsedDSN.Query(); query.Get("parseTime") == "" {
		query.Set("parseTime", "true")
		parsedDSN.RawQuery = query.Encode()
	}

	db.dsn = parsedDSN.String()

	db.DB, err = gorm.Open(mysql.Open(db.dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
