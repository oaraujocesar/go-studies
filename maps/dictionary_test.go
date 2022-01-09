package dictionary

import "testing"

func TestSearch(t *testing.T) {

	assertStrings := func(t testing.TB, received, expected, word string) {
		t.Helper()

		if received != expected {
			t.Errorf("Received %q, but expected %q given %q", received, expected, word)
		}
	}

	dictionary := map[string]string{"test": "this is just a test"}

	word := "test"
	received := Search(dictionary, word)
	expected := "this is just a test"

	assertStrings(t, received, expected, word)
}
