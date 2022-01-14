package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	received := buffer.String()
	expected := "Hello, Chris!"

	if received != expected {
		t.Errorf("Received %q, but expected %q", received, expected)
	}
}
