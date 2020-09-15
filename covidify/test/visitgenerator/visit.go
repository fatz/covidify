package visitgenerator

import (
	"math/rand"

	"github.com/fatz/covidify/covidify/models"
	"syreclabs.com/go/faker"
)

const MaxTableVisitors int = 6

func NewFakeVisit(tablenr string, tables []string) *models.Visit {
	table := "foo"

	if len(tables) >= 1 {
		table = tables[rand.Intn(len(tables))]
	}

	v := models.NewVisit()
	v.TableNumber = table

	v.Visitors = make([]models.Visitor, rand.Intn(MaxTableVisitors))

	for i, _ := range v.Visitors {
		v.Visitors[i] = *NewFakeVisitor()
	}

	return &v
}

func NewFakeVisitor() *models.Visitor {
	v := new(models.Visitor)

	addr := faker.Address()
	v.City = addr.City()
	v.Country = addr.Country()
	v.Street = addr.StreetAddress()
	v.ZipCode = addr.ZipCode()

	v.Phone = faker.PhoneNumber().String()
	v.Email = faker.Internet().Email()

	name := faker.Name()
	v.Name = name.Name()

	return v
}
