package maps

import "testing"

func assertDict(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%q wanted but got %q this", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("expected an error %q but got %q", want, got)
	}
}

func TestSearch(t *testing.T) {

	t.Run("search a key that exists in dict", func(t *testing.T) {
		d := Dictionary{"Ant": "Ant is an insect"}

		got, _ := d.Search("Ant")
		want := "Ant is an insect"
		assertDict(t, got, want)

	})

	t.Run("search a key that dont exist in dict", func(t *testing.T) {
		dict := Dictionary{"Bat": "bat is a mammal"}
		_, got := dict.Search("Ant")
		assertError(t, got, ErrKeyNotFound)
	})

}

func TestAddWord(t *testing.T) {

	assertDict := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("%q wanted but got %q this", want, got)
		}
	}

	t.Run("Add a word to dictionary", func(t *testing.T) {
		dict := Dictionary{}
		word := "Bat"
		description := "Bat is a mammal"
		ok := dict.AddWord(word, description)

		got, err := dict.Search(word)

		assertError(t, err, ok)
		assertDict(t, got, description)
	})

	t.Run("Trying to add an existing word", func(t *testing.T) {
		dict := Dictionary{"Cat": "cat is cute"}
		err := dict.AddWord("Cat", "cat is not so cute")
		assertError(t, err, ErrKeyAlreadyExists)
		assertDict(t, dict["Cat"], dict["Cat"])

	})

}

func TestUpdateDescription(t *testing.T) {
	t.Run("update the description", func(t *testing.T) {
		dict := Dictionary{"Cat": "cat is cute"}
		updated_description := "cat is not so cute"
		err := dict.UpdateDescription("Cat", updated_description)
		assertError(t, err, nil)
		assertDict(t, dict["Cat"], updated_description)
	})

	t.Run("update description if word doesnt exist", func(t *testing.T) {
		dict := Dictionary{"Bat": "bat is mammal"}
		updated_description := "cat is so cute"
		err := dict.UpdateDescription("Cat", updated_description)
		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDeleteWord(t *testing.T) {
	t.Run("delete a existing word", func(t *testing.T) {
		dict := Dictionary{"Bat": "bat is mammal"}
		err := dict.DeleteWord("Bat")
		assertError(t, err, nil)
		assertDict(t, dict["Bat"], "")
	})

	t.Run("delete an word which dont exist", func(t *testing.T) {
		dict := Dictionary{"Bat": "bat is mammal"}
		err := dict.DeleteWord("Cat")
		assertError(t, err, ErrWordDoesNotExists)
	})
}
