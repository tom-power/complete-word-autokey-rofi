package completeword

import (
	"fmt"
	"os/user"
)

const wordsPath = "/.config/complete-word-autokey-rofi/words/"

// Complete a selection
func Complete(
  getWords GetWords,
  chooseWord ChooseWord,
  selection string) {
	usr, _ := user.Current()
	home := usr.HomeDir
	words, err := getWords(home + wordsPath)
	if err != nil {
		fmt.Printf(err.Error())
	}
	completion, err := chooseWord(selection, words)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf(completion)
}
