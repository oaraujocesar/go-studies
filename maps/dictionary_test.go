package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	word := "test"
	received := Search(dictionary, word)
	expected := "this is just a test"

	if received != expected {
		t.Errorf("Received %q, but expected %q given %q", received, expected, word)
	}
}
