package stopwatch

import (
	"reflect"
	"testing"

	"fmt"
	"time"
)

///////////////////////////////////////////////////////////////////////
//// Examples
///////////////////////////////////////////////////////////////////////
func Example() {
	  sw := Start()
	  time.Sleep(1 * time.Second)
    lap1_time := sw.Elapsed()
	  time.Sleep(1 * time.Second)
    stop_time := sw.Stop()
	  time.Sleep(1 * time.Second)
    // Stopwatch no longer increments after Stop is called.
    // So this call to Elapsed() returns the same thing Stop() did.
    late_time := sw.Elapsed()

    fmt.Println(int(lap1_time))
    fmt.Println(int(stop_time))
    fmt.Println(int(late_time))
    // Output:
    // 1
    // 2
    // 2
}

///////////////////////////////////////////////////////////////////////
//// Unit Tests
///////////////////////////////////////////////////////////////////////
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
		t.Fatal("Elapsed() is not incrementing after Start().")
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