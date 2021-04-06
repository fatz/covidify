package db

import (
	"time"

	models "github.com/fatz/covidify/covidify/models"
)

// CreateVisit inserts Visit into DB
func (d *DB) CreateVisit(v models.Visit) (*models.Visit, error) {
	if err := v.Valid(); err != nil {
		return nil, err
	}

	res := d.DB.Create(&v)
	if res.Error != nil {
		return nil, res.Error
	}

	return &v, nil
}

func (d *DB) GetVisit(id string) (*models.Visit, error) {
	var v models.Visit

	res := d.DB.Preload("Visitors").First(&v, "`visits`.`id`= ?", id)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected < 1 {
		return nil, nil
	}

	return &v, nil
}

func (d *DB) GetTables() ([]string, error) {
	tables := make([]string, 0)

	res := d.DB.Model(&models.Visit{}).Distinct().Pluck("table_number", &tables)
	if res.Error != nil {
		return nil, res.Error
	}

	return tables, nil
}

func (d *DB) GetVisitsByTable(tableNumber string) ([]models.Visit, error) {
	var visits []models.Visit

	res := d.DB.Where("table_number = ?", tableNumber).Find(&visits)
	if res.Error != nil {
		return nil, res.Error
	}

	return visits, nil
}

func (d *DB) GetVisitsByTableCheckinBetweeen(tableNumber string, after, before time.Time) ([]models.Visit, error) {
	var visits []models.Visit

	res := d.DB.Where("table_number = ? AND checkin > ? AND checkin < ?", tableNumber, after, before).Find(&visits)
	if res.Error != nil {
		return nil, res.Error
	}

	return visits, nil
}

func (d *DB) DeleteVisitsByTableCheckinBetweeen(tableNumber string, after, before time.Time) error {
	res := d.DB.Where("table_number = ? AND checkin > ? AND checkin < ?").Delete(&models.Visit{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
