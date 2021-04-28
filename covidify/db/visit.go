package db

import (
	"errors"
	"time"

	models "github.com/fatz/covidify/covidify/models"
	"gorm.io/gorm"
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

// GetVisit returns a given `visit`
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

// GetTables returns all tables with visits
func (d *DB) GetTables() ([]string, error) {
	tables := make([]string, 0)

	res := d.DB.Model(&models.Visit{}).Distinct().Pluck("table_number", &tables)
	if res.Error != nil {
		return nil, res.Error
	}

	return tables, nil
}

// GetVisitsByTable returns all visits for a given `tableNumber`
func (d *DB) GetVisitsByTable(tableNumber string) ([]models.Visit, error) {
	var visits []models.Visit

	res := d.DB.Where("table_number = ?", tableNumber).Find(&visits)
	if res.Error != nil {
		return nil, res.Error
	}

	return visits, nil
}

// GetVisitsByTableCheckinBetweeen gets all visists of a given `tableNumber` betweeen `after` and `before`
func (d *DB) GetVisitsByTableCheckinBetweeen(tableNumber string, after, before time.Time) ([]models.Visit, error) {
	var visits []models.Visit

	res := d.DB.Where("table_number = ? AND check_in > ? AND check_in < ?", tableNumber, after, before).Find(&visits)
	if res.Error != nil {
		return nil, res.Error
	}

	return visits, nil
}

// GetVisitsByTableCheckinBetweeen gets all visists of a given `tableNumber` betweeen `after` and `before`
func (d *DB) GetVisitsByTableCheckinBetweeenFirst(tableNumber string, after, before time.Time) (*models.Visit, error) {
	var visit models.Visit

	res := d.DB.Where("table_number = ? AND check_in > ? AND check_in < ?", tableNumber, after, before).First(&visit)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}

	return &visit, nil
}

// GetVisitCheckinBetweeenFirst gets the first visist betweeen `after` and `before`
func (d *DB) GetVisitCheckinBetweeenFirst(after, before time.Time) (*models.Visit, error) {
	var visit models.Visit

	res := d.DB.Where("check_in > ? AND check_in < ?", after, before).First(&visit)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}

	return &visit, nil
}

// GetVisitCheckinBetweeenFirst gets the first visist betweeen `after` and `before`
func (d *DB) GetVisitCheckinBetweeenLimit(after, before time.Time, limit int) ([]models.Visit, error) {
	var visits []models.Visit

	res := d.DB.Where("check_in > ? AND check_in < ?", after, before).Limit(limit).Find(&visits)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}

	return visits, nil
}

// DeleteVisit deletes a single Visit
func (d *DB) DeleteVisit(visit *models.Visit) error {
	res := d.DB.Delete(visit)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// DeleteVisit deletes a single Visit
func (d *DB) DeleteVisitorsByVisitID(visitID string) (int64, error) {
	var visitor models.Visitor

	res := d.DB.Delete(visitor, "visit_id = ?", visitID)
	if res.Error != nil {
		return res.RowsAffected, res.Error
	}

	return res.RowsAffected, nil
}

// DeleteVisitsCheckinBetweeen cleans up all visists betweeen `after` and `before`
func (d *DB) DeleteVisitsCheckinBetweeen(after, before time.Time) error {
	res := d.DB.Where("check_in > ? AND check_in < ?", after, before).Delete(&models.Visit{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
