package dictionary

import "testing"

func TestSearch(t *testing.T) {

	assertStrings := func(t testing.TB, received, expected, word string) {
		t.Helper()

		if received != expected {
			t.Errorf("Received %q, but expected %q given %q", received, expected, word)
		}
	}

	assertError := func(t testing.TB, received, expected error, word string) {
		t.Helper()

		if received != expected {
			t.Errorf("Received error %q, but expected %q", received, expected)
		}
	}

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("it should return the definition if a known word is passed", func(t *testing.T) {
		word := "test"

		received, _ := dictionary.Search(word)
		expected := "this is just a test"

		assertStrings(t, received, expected, word)
	})

	t.Run("it should return an error if an unknown word is passed", func(t *testing.T) {
		_, err := dictionary.Search("hey")

		assertError(t, err, ErrNotFound, "hey")
	})

}
