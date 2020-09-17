package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFakeVisit(t *testing.T) {
	table := "foo1-123"

	visit := NewFakeVisit(table)

	if assert.NotNil(t, visit) {
		assert.Equal(t, table, visit.TableNumber)
		if assert.GreaterOrEqual(t, 1, len(visit.Visitors)) {
			assert.NotEmpty(t, visit.Visitors[0].City)
			assert.NotEmpty(t, visit.Visitors[0].Country)
			assert.NotEmpty(t, visit.Visitors[0].Email)
			assert.NotEmpty(t, visit.Visitors[0].Name)
			assert.NotEmpty(t, visit.Visitors[0].ZipCode)
			assert.NotEmpty(t, visit.Visitors[0].Street)
			assert.NotEmpty(t, visit.Visitors[0].Phone)
		}
	}

}
