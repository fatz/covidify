package db

import (
	models "github.com/fatz/covidify/covidify/models"
	"github.com/relops/cqlr"
)

// CreateVisit Insters Visit into DB
func (d *DB) CreateVisit(v models.Visit) (*models.Visit, error) {
	if err := v.Valid(); err != nil {
		return nil, err
	}

	sess, err := d.Session()
	if err != nil {
		return nil, err
	}

	q := cqlr.Bind("INSERT INTO visit (id, checkin, table_number, visitors) VALUES (?, ?, ?, ?)", v)

	err = q.Exec(sess)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func (d *DB) GetVisit(id string) (*models.Visit, error) {
	var v models.Visit
	sess, err := d.Session()
	if err != nil {
		return nil, err
	}

	q := sess.Query("SELECT * from visit WHERE id = ?", id)
	b := cqlr.BindQuery(q)

	b.Scan(&v)

	return &v, nil
}
