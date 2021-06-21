package completeword

func Complete(
	getWords GetWords,
	chooseWord ChooseWord,
	selection string) (string, error) {
	words, err := getWords()
	if err != nil {
		return "", err
	}
	return chooseWord(words, selection)
}
