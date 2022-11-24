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

func TestElapsed(t *testing.T) {
	s := Start()
	time.Sleep(300 * time.Millisecond)
	elapsed_val := s.Elapsed()
	s.Stop()

	if elapsed_val == 0 {
		t.Fatal("Elapsed() is not increasing after stopwatch is started.")
	}	
}

func TestElapsed_Increases(t *testing.T) {
	s := Start()
	time.Sleep(300 * time.Millisecond)
	first_check := s.Elapsed()
	time.Sleep(300 * time.Millisecond)
	second_check := s.Elapsed()
	s.Stop()

	if second_check <= first_check {
		t.Fatal("Elapsed() is not incrementing in subsequent calls after start.")
	}
}

func TestElapsed_AfterStop(t *testing.T) {
	s := Start()
	time.Sleep(300 * time.Millisecond)
	stopped_val := s.Stop()
	elapsed_val := s.Elapsed()
	if elapsed_val > stopped_val {
		t.Fatal("Elapsed() is not stopping after Stop() is called.")
	}
	if elapsed_val != stopped_val {
		t.Fatal("Elapsed() is not returning same value as Stop().")
	}
}

// Examples