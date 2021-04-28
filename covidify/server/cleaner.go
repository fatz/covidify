package covidify

import (
	"time"

	"github.com/prometheus/common/log"
)

const Limit = 1000

func (s *Server) Clean(olderThan time.Time) error {
	tables, err := s.db.GetTables()
	if err != nil {
		return err
	}

	s.config.Logger.Tracef("Found %d tables to cleanup", len(tables))
	s.statsDIncrementByValue("covidify.clean.tables", len(tables))

	var visitorCounter int64

	for {
		// TODO: we expect this to be slow but not blocking. Should be improved to select for update
		visits, err := s.db.GetVisitCheckinBetweeenLimit(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC), olderThan, Limit)
		if err != nil {
			s.config.Logger.Warnf("Got error while requesting visit %v", err)
			continue
		}
		lenVisits := len(visits)

		for _, visit := range visits {
			s.config.Logger.Debugf("Cleaning visit %s", visit.Id)

			numVisitors, err := s.db.DeleteVisitorsByVisitID(visit.Id)
			visitorCounter = visitorCounter + numVisitors

			if err != nil {
				s.config.Logger.Warnf("Error cleaning up Visitors for visit %s - %v", visit.Id, err)
			}

			if err := s.db.DeleteVisit(&visit); err != nil {
				s.config.Logger.Errorf("Could not clean Visit %s - %v", visit.Id, err)
				continue
			}
		}
		s.config.Logger.Infof("Cleaned %d visits with %d visitors. Reseting Counter", lenVisits, visitorCounter)
		visitorCounter = 0

		if lenVisits < 1000 {
			log.Infof("Found less records (%d) than batch size (%d) - ending clean process", lenVisits, Limit)
			return nil
		}
	}

	return nil
}
