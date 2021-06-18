package completeword

// Complete a selection
func Complete(
  getWords GetWords,
  chooseWord ChooseWord,
  selection string) (string, error) {
	words, err := getWords()
	if err != nil {
		return "", err
	}
	return chooseWord(selection, words)
}
