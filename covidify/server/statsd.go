package covidify

import "strings"

func dropEmpty(in []string) (o []string) {
	for _, s := range in {

		if st := strings.TrimSpace(s); st != "" {
			o = append(o, st)
		}
	}

	return o
}

func (s *Server) statsDPrefixStat(stat string) string {
	p := strings.Split(s.config.StatsDPrefix, ".")
	p = append(p, strings.Split(stat, ".")...)
	return strings.Join(dropEmpty(p), ".")
}

// StatsDIncrement wraps statsd.Increment() adding prefix
func (s *Server) statsDIncrement(stat string) {
	if s.statsd == nil {
		return
	}
	s.statsd.Increment(s.statsDPrefixStat(stat))
}

// statsDIncrementByValue wraps statsd.IncrementByValue() adding prefix
func (s *Server) statsDIncrementByValue(stat string, val int) {
	if s.statsd == nil {
		return
	}
	s.statsd.IncrementByValue(s.statsDPrefixStat(stat), val)
}

// statsDIncrementWithSampling wraps statsd.IncrementWithSampling() adding prefix
func (s *Server) statsDIncrementWithSampling(stat string, sampleRate float32) {
	if s.statsd == nil {
		return
	}
	s.statsd.IncrementWithSampling(s.statsDPrefixStat(stat), sampleRate)
}

// statsDDecrement wraps statsd.Decrement() adding prefix
func (s *Server) statsDDecrement(stat string) {
	if s.statsd == nil {
		return
	}
	s.statsd.Decrement(s.statsDPrefixStat(stat))
}

// statsDDecrementWithSampling wraps statsd.DecrementWithSampling() adding prefix
func (s *Server) statsDDecrementWithSampling(stat string, sampleRate float32) {
	if s.statsd == nil {
		return
	}
	s.statsd.DecrementWithSampling(s.statsDPrefixStat(stat), sampleRate)
}
