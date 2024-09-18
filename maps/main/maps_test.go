package main

import "testing"

func TestMaps(t *testing.T) {
	verifyString := func(t *testing.T, result, expect string) {
		t.Helper()
		if result != expect {
			t.Errorf("result %s, want %s", result, expect)
		}
	}

	verifyError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Run("search for known word", func(t *testing.T) {
		dictionary := MyDictionary{"Jane": "Jane Due"}
		result, _ := dictionary.Search("Jane")
		expected := "Jane Due"
		verifyString(t, result, expected)
	})

	t.Run("search for unknown word", func(t *testing.T) {
		dictionary := MyDictionary{"Jane": "Jane Due"}
		_, err := dictionary.Search("Jane")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("add itens for map", func(t *testing.T) {
		dictionary := MyDictionary{}
		dictionary.Add("Jane", "Jane Due")

		result, err := dictionary.Search("Jane")
		expected := "Jane Due"

		verifyError(t, err)
		verifyString(t, result, expected)
	})

	t.Run("add items to map word already exists", func(t *testing.T) {
		dictionary := MyDictionary{"Jane": "Jane Due"}
		err := dictionary.Add("Jane", "Jane Due")

		verifyError(t, err)
	})

	t.Run("update map", func(t *testing.T) {
		dictionary := MyDictionary{"Jane": "Jane Due"}
		err := dictionary.Update("Jane", "Jane for")

		verifyError(t, err)
	})

	t.Run("delete map", func(t *testing.T) {
		dictionary := MyDictionary{"Jane": "Jane Due"}
		err := dictionary.Delete("Jane")
		verifyError(t, err)
	})
}
