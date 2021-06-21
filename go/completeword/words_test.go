package completeword

import (
	"strings"
	"testing"
)

func Test_readWords(t *testing.T) {
	t.Run("can get words", func(t *testing.T) {
		testWords := "one\nword\nper\nline"
		words, err := readWords(strings.NewReader(testWords))
		if err != nil {
			t.Errorf("error = %v", err)
		}
		if strings.Join(words, "\n") != testWords {
			t.Errorf("error = %v", words)
		}
	})
}
