package visitgenerator

import (
	"fmt"
	"math/rand"
)

func GenTableList(restaurants, maxTables int) []string {
	s := make([]string, 0)

	for i := 0; i < restaurants; i++ {
		s = append(s, GenRestaurant(fmt.Sprintf("rest%02d", i), maxTables)...)
	}
	return s
}

func GenRestaurant(restaurant string, maxTables int) []string {
	s := make([]string, rand.Intn(maxTables))

	for i, _ := range s {
		s[i] = fmt.Sprintf("%s-%05d", restaurant, i)
	}
	return s
}
