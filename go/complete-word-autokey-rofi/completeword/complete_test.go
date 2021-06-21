package completeword

import (
	"testing"
)

var chooseFirst = func(
  selection string, 
  words []string) (string, error) {
    return words[0], nil
  }
  
var testWords = func() ([]string, error) { 
	return []string{"test","testing"}, nil
}

func Test_complete(t *testing.T) {
	t.Run("can complete", func(t *testing.T) {
	  completeWord, err := Complete(testWords, chooseFirst, "te")
		if err != nil {
			t.Errorf("error = %v", err)
		}
		if completeWord != "test" {
			t.Errorf("error = %v", completeWord)
		}
	})
}
