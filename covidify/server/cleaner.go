package covidify

import (
	"fmt"
	"strings"
	"time"
)

func (s *Server) Clean(olderThan time.Time) error {
	tables, err := s.db.GetTables()
	if err != nil {
		return err
	}

	s.config.Logger.Tracef("Found %d tables to cleanup", len(tables))

	errors := make([]string, 0)
	for _, t := range tables {
		if err := s.CleanTable(t, olderThan); err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("Found %d errors - %s", len(errors), strings.Join(errors, ","))
	} else {
		return nil
	}
}

func (s *Server) CleanTable(tableNumber string, olderThan time.Time) error {
	s.config.Logger.Tracef("Cleaning Up visits older then %s on table %s", olderThan.Format(time.RFC1123), tableNumber)
	return s.db.DeleteVisitsByTableCheckinBetweeen(tableNumber, time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), olderThan)
}
