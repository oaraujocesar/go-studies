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

func TestAdd(t *testing.T) {
	t.Run("it should add a new word and its definition to the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "hey"
		definition := "a simple greeting"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("it should throw an error if the word already exists", func(t *testing.T) {
		word := "hello"
		definition := "the simplest way of greeting"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertError(t testing.TB, received, expected error) {
	t.Helper()
	if received != expected {
		t.Errorf("received %q expected %q", received, expected)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	received, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != received {
		t.Errorf("received %q want %q", received, definition)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("it should update an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		newDefinition := "this is the new definition for test"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("it should return an error if word does not exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}
