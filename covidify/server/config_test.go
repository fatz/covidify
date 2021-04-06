package covidify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDBDSN(t *testing.T) {
	db := NewConfig()

	assert.NotNil(t, db)
	assert.Equal(t, "mysql://tcp(127.0.0.1:3306)/covidify", db.GenDBDSN())
}
