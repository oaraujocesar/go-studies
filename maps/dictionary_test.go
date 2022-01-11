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

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	received, err := dictionary.Search("hey")

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	if received != definition {
		t.Errorf("Received %q, but expected %q given %q", received, definition, word)
	}
}

func TestAdd(t *testing.T) {
	t.Run("it should add a new word and its definition to the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "hey"
		definition := "a simple greeting"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})
}
