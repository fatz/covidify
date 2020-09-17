package main

import (
	"math/rand"

	"github.com/fatz/covidify/covidify/models"
	"syreclabs.com/go/faker"
)

const MaxTableVisitors int = 6

func NewFakeVisit(table string) *models.Visit {
	v := models.NewVisit()
	v.TableNumber = table
	numVisitors := 1

	if randNum := rand.Intn(MaxTableVisitors); randNum > 1 {
		numVisitors = randNum
	}

	v.Visitors = make([]models.Visitor, numVisitors)

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
