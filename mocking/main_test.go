package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)

	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("it should print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		Countdown(buffer, spySleeper)

		received := buffer.String()
		expected := `3
2
1
Go!`

		if received != expected {
			t.Errorf("Received %q, but expected %q", received, expected)
		}

	})

	t.Run("it should sleep before every print", func(t *testing.T) {
		SpySleepPrinter := &SpyCountdownOperations{}

		Countdown(SpySleepPrinter, SpySleepPrinter)

		expected := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expected, SpySleepPrinter.Calls) {
			t.Errorf("Expected calls%v but received %v", expected, SpySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
