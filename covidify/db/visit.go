package db

import (
	"fmt"
	"time"

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

	q := cqlr.Bind("INSERT INTO "+d.Keyspace+".visit (id, checkin, table_number, visitors) VALUES (?, ?, ?, ?)", v)

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

	q := sess.Query("SELECT * from "+d.Keyspace+".visit WHERE id = ?", id)
	b := cqlr.BindQuery(q)

	b.Scan(&v)

	return &v, nil
}

func (d *DB) GetTables() ([]string, error) {
	sess, err := d.Session()
	if err != nil {
		return nil, err
	}

	s := make([]string, 0)

	q := sess.Query("SELECT DISTINCT table_number FROM " + d.Keyspace + ".visit GROUP BY table_number")
	iter := q.Iter()

	var tn string

	for iter.Scan(&tn) {
		s = append(s, tn)
	}

	return s, nil
}

func (d *DB) GetVisitsByTable(tableNumber string) ([]models.Visit, error) {
	sess, err := d.Session()
	if err != nil {
		return nil, err
	}

	v := make([]models.Visit, 0)

	q := sess.Query("SELECT * from "+d.Keyspace+".visit WHERE table_number = ?", tableNumber)
	b := cqlr.BindQuery(q)

	b.Scan(&v)

	return v, err
}

func (d *DB) GetVisitsByTableCheckinBetweeen(tableNumber string, after, before time.Time) ([]models.Visit, error) {
	sess, err := d.Session()
	if err != nil {
		return nil, err
	}

	v := make([]models.Visit, 0)

	q := sess.Query("SELECT * from "+d.Keyspace+".visit WHERE table_number = ? AND checkin > ? AND checkin < ?", tableNumber, after, before)
	b := cqlr.BindQuery(q)

	b.Scan(&v)

	return v, err
}

func (d *DB) DeleteVisitsByTableCheckinBetweeen(tableNumber string, after, before time.Time) error {
	sess, err := d.Session()
	if err != nil {
		return err
	}

	q := sess.Query("DELETE from "+d.Keyspace+".visit WHERE table_number = ? AND checkin > ? AND checkin < ?", tableNumber, after, before)
	fmt.Println(q.String())
	err = q.Exec()
	if err != nil {
		return err
	}

	return err
}
