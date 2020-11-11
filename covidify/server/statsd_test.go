package covidify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestdropEmpty(t *testing.T) {
	assert.Equal(t, []string{"foo", "bar"}, dropEmpty([]string{" foo", " ", "", "bar "}))
}
