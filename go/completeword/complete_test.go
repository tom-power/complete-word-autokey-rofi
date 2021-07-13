package completeword

import (
	"errors"
	"testing"
)

var chooseMatching ChooseWord = func(words []string, selection string) (string, error) {
	for _, word := range words {
		if word == selection {
			return word, nil
		}
	}
	return "", errors.New("no match")
}

var testWords GetWords = func() ([]string, error) {
	return []string{"foo", "bar", "baz"}, nil
}

func Test_complete(t *testing.T) {
	t.Run("can complete", func(t *testing.T) {
		completeWord, err := Complete(testWords, chooseMatching)("foo")
		if err != nil {
			t.Errorf("error = %v", err)
		}
		if completeWord != "foo" {
			t.Errorf("error = %v", completeWord)
		}
	})

	t.Run("keep titlecase from selection", func(t *testing.T) {
		completeWord, err := Complete(testWords, chooseMatching)("Foo")
		if err != nil {
			t.Errorf("error = %v", err)
		}
		if completeWord != "Foo" {
			t.Errorf("error = %v", completeWord)
		}
	})
}
