package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	received := buffer.String()
	expected := `3
2
1
Go!`

	if received != expected {
		t.Errorf("Received %q, but expected %q", received, expected)
	}
}
