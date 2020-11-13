package main

import (
	"math/rand"
	"time"

	"github.com/fatz/covidify/covidify/models"
	"syreclabs.com/go/faker"
)

const MaxTableVisitors int = 6
const DaysBack int = 10
const MaxStayMinutes int = 180
const MinStayMinutes int = 10

func NewFakeVisit(table string) *models.Visit {
	v := models.NewVisit()
	v.TableNumber = table
	numVisitors := 1

	v.CheckIn = faker.Date().Backward(time.Duration(DaysBack) * 24 * time.Hour)

	minCheckout := v.CheckIn.Add(time.Duration(MinStayMinutes) * time.Minute)
	maxCheckout := v.CheckIn.Add(time.Duration(MaxStayMinutes) * time.Minute)
	v.CheckOut = faker.Time().Between(minCheckout, maxCheckout)

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
