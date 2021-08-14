package completeword_test

import (
	"../completeword"
	"testing"
)

func Test_add(t *testing.T) {

	var chooseFirst completeword.ChooseWord = func(words []string, selection string) (string, error) {
		return words[0], nil
	}

	var chooseSelection completeword.ChooseWord = func(words []string, selection string) (string, error) {
		return selection, nil
	}

	var testWords completeword.GetWords = func() ([]string, error) {
		return []string{"test", "testing"}, nil
	}

	var noPath completeword.GetPath = func() (string, error) {
		return "", nil
	}

	var noopWriteWord completeword.WriteWord = func(path string, word string) error {
		return nil
	}

	t.Run("can add", func(t *testing.T) {
		err := completeword.Add(testWords, chooseSelection, noPath, noopWriteWord)("hello")
		if err != nil {
			t.Errorf("error = %v", err)
		}
	})

	t.Run("can't add duplicate", func(t *testing.T) {
		err := completeword.Add(testWords, chooseFirst, noPath, noopWriteWord)("te")
		if err == nil {
			t.Error("expecting an error")
		}
	})
}
