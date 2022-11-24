package stopwatch

import (
	"reflect"
	"testing"

	"time"
)

func TestStart(t *testing.T) {
	s := Start()
	if reflect.TypeOf(s).String() != "stopwatch.Stopwatch" {
	  t.Fatal("Start() did not return a Stopwatch type.")
  }
}

func TestStop(t *testing.T) {
	s := Start()
	time.Sleep(300 * time.Millisecond)
	time_at_stop := s.Stop()
	time.Sleep(300 * time.Millisecond)
	if s.Elapsed() > time_at_stop {
		t.Fatal("Stop() did not stop timer.")
	}
}

// Test for Elapsed returning increasing values as time passes.

// Test for Elapsed returning same value after Stop is called.

// Examples