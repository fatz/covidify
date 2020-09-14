package db

import (
	"testing"

	"github.com/stretchr/testify/assert"

	covidify "github.com/fatz/covidify/covidify/server"
)

func testCreateVisit() covidify.Visit {

	v := covidify.NewVisit()
	v.TableNumber = "o-123"
	vs := covidify.Visitor{}
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

	d, err := envConnect()

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

	d, err := envConnect()

	assert.NoError(t, err)
	assert.NotNil(t, d)

	v1 := testCreateVisit()

	_, err = d.CreateVisit(v1)
	if assert.NoError(t, err) {
		v2, err := d.GetVisit(v1.Id)
		assert.NoError(t, err)
		assert.Equal(t, v1.TableNumber, v2.TableNumber)
		assert.Equal(t, v1.Visitors, v2.Visitors)
	}

}
