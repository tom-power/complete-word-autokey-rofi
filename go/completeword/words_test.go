package completeword

import (
	"strings"
	"testing"
)

func Test_readWords(t *testing.T) {
	t.Run("can get words", func(t *testing.T) {
		words, err := readWords(strings.NewReader("one\nword\nper\nline"))
		if err != nil {
			t.Errorf("error = %v", err)
		}
		if words != "one\nword\nper\nline" {
			t.Errorf("error = %v", words)
		}
	})
}
