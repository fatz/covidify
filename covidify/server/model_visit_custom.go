package covidify

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewVisit() Visit {
	v := Visit{}

	u, err := uuid.NewRandom()
	if err == nil {
		v.Id = u.String()
	}

	v.CheckIn = time.Now()

	v.Visitors = make([]Visitor, 0)

	return v
}

func (v *Visit) Valid() error {
	var errors []string

	if v.Id == "" {
		errors = append(errors, "Id missing")
	}

	if v.TableNumber == "" {
		errors = append(errors, "TableNumber missing")
	}

	if v.CheckIn.Unix() <= 0 {
		errors = append(errors, "CheckIn invalid")
	}

	if len(v.Visitors) < 1 {
		errors = append(errors, "Visitors must contain at least 1 Visitor")
	}

	if e := len(errors); e > 0 {
		return fmt.Errorf("Found %d errors: %s", e, strings.Join(errors, ","))
	}

	return nil
}
