package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/fatz/covidify/covidify/models"
)

func testCreateVisit() models.Visit {

	v := models.NewVisit()
	v.TableNumber = "o-123"
	vs := models.Visitor{}
	vs.City = "Hamburg"
	vs.Country = "DEU"
	vs.Name = "Max Mustermann"
	vs.Phone = "+4940123456789"
	vs.Street = "Am Sandtorkai 43"
	vs.ZipCode = "20457"
	v.Visitors = append(v.Visitors, vs)

	return v
}

func TestCreateVisitIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	d, err := NewDB(dsn)

	assert.NoError(t, err)
	assert.NotNil(t, d)

	v := testCreateVisit()

	_, err = d.CreateVisit(v)
	assert.NoError(t, err)

}

func TestGetVisitIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	d, err := NewDB(dsn)

	if assert.NoError(t, err) && assert.NotNil(t, d) {

		v1 := testCreateVisit()

		res1, err := d.CreateVisit(v1)
		assert.NotNil(t, res1)

		if assert.NoError(t, err) {
			v2, err := d.GetVisit(res1.Id)
			if assert.NoError(t, err) && assert.NotNil(t, v2) {
				assert.Equal(t, v1.TableNumber, v2.TableNumber)
				assert.Equal(t, v1.Visitors, v2.Visitors)
			}
		}
	}

}

func TestGetVisitsByTableCheckinBetweeen(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	vis1 := testCreateVisit()
	vis1.CheckIn = time.Date(2020, 9, 1, 21, 0, 0, 0, time.UTC)
	vis1Checkout := time.Date(2020, 9, 1, 21, 30, 0, 0, time.UTC)
	vis1.CheckOut = &vis1Checkout

	vis2 := testCreateVisit()
	vis2.CheckIn = time.Date(2020, 9, 2, 21, 0, 0, 0, time.UTC)
	vis2Checkout := time.Date(2020, 9, 2, 21, 30, 0, 0, time.UTC)
	vis2.CheckOut = &vis2Checkout

	vis3 := testCreateVisit()
	vis3.CheckIn = time.Date(2020, 10, 1, 21, 0, 0, 0, time.UTC)
	vis3Checkout := time.Date(2020, 10, 1, 21, 30, 0, 0, time.UTC)
	vis3.CheckOut = &vis3Checkout

	d, err := NewDB(dsn)

	if assert.NoError(t, err) && assert.NotNil(t, d) {
		d.CreateVisit(vis1)
		d.CreateVisit(vis2)
		d.CreateVisit(vis3)

		visits, err := d.GetVisitsByTableCheckinBetweeen(vis1.TableNumber, vis1.CheckIn.Add(-time.Hour), vis2.CheckOut.Add(time.Hour))
		if assert.NoError(t, err) {
			assert.Len(t, visits, 2)
		}

	}

	// also test Delete
	if err := d.DeleteVisitsCheckinBetweeen(vis1.CheckIn.Add(-time.Hour), vis2.CheckOut.Add(time.Hour)); assert.NoError(t, err) {
		visits, err := d.GetVisitsByTableCheckinBetweeen(vis1.TableNumber, vis1.CheckIn.Add(-time.Hour), vis2.CheckOut.Add(time.Hour))
		if assert.NoError(t, err) {
			assert.Len(t, visits, 0)
		}
	}

}
