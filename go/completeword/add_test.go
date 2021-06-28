package completeword

import (
	"testing"
)

func Test_add(t *testing.T) {

  var chooseFirst ChooseWord = func(words []string, selection string) (string, error) {
  	return words[0], nil
  }

  var chooseSelection ChooseWord = func(words []string, selection string) (string, error) {
  	return selection, nil
  }

  var testWords GetWords = func() ([]string, error) {
  	return []string{"test", "testing"}, nil
  }

  var noPath GetPath = func() (string, error) {
    return "", nil
  }

  var noopWriteWord WriteWord = func(path string, word string) error {
  	return nil
  }

	t.Run("can add", func(t *testing.T) {
		err := Add(testWords, chooseSelection, "hello", noPath, noopWriteWord)
		if err != nil {
			t.Errorf("error = %v", err)
		}
	})

	t.Run("can't add duplicate", func(t *testing.T) {
    err := Add(testWords, chooseFirst, "te", noPath, noopWriteWord)
		if err == nil {
			t.Error("expecting an error")
		}
	})
}
