package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "word not found"

		if err == nil {
			t.Fatal("we expect an error but got none")
		}

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add non existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		_ = dictionary.Add("test", "this is a test")

		want := "this is a test"
		got, err := dictionary.Search("test")
		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("error on existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		err := dictionary.Add("test", "this is a test")

		assertStrings(t, err.Error(), ErrWordExists.Error())
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("we got an error when we expected none")
	}
}