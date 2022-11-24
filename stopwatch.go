// Package stopwatch is a simple library to help time how long things take.
// It can be very useful for timing scripts or providing runtimes to users.
package stopwatch

import "time"

// Stopwatch is the core struct used. It should not be created directly,
// but instead through the use for the New() function.
// No fields are exported, data is retrieved via method calls.
type Stopwatch struct {
	started time.Time
	stopped time.Time
}

// Start is how a new instance of Stopwatch is created. The returned
// Stopwatch starts "ticking" when this function is called.
func Start() (sw Stopwatch) {
	sw.started = time.Now()
	return
}

// Elapsed returns time.Now()-[start time] if stopwatch is still running.
// Will return [stopped time] - [started time] if stopwatch has been stopped.
func (s *Stopwatch) Elapsed() float64 {
	if (s.stopped != time.Time{}) {
		return s.stopped.Sub(s.started).Seconds()
	}
	return time.Since(s.started).Seconds()
}

// Stop will stop the stopwatch and return the total elapsed time.
func (s *Stopwatch) Stop() float64 {
	s.stopped = time.Now()
	return s.Elapsed()
}