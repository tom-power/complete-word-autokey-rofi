package completeword

import (
	"strings"
	"unicode"
)

func Complete(
	getWords GetWords,
	chooseWord ChooseWord,
) func(string) (string, error) {
	return func(selection string) (string, error) {
		words, err := getWords()
		if err != nil {
			return "", err
		}
		completion, err := chooseWord(words, strings.ToLower(selection))
		if err != nil {
			return "", err
		}
		return keepTitleCase(selection, completion), nil
	}
}

func keepTitleCase(selection string, completion string) string {
	if unicode.IsUpper([]rune(selection)[0]) {
		return strings.Title(completion)
	}
	return completion
}
