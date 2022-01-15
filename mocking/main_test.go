package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	received := buffer.String()
	expected := `3
2
1
Go!`

	if received != expected {
		t.Errorf("Received %q, but expected %q", received, expected)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enought calls to sleeper, expected 4 but got it %d", spySleeper.Calls)
	}
}
