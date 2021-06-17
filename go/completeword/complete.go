package completeword

import (
	"fmt"
	"os/user"
)

const wordsPath = "/.config/complete-word-autokey-rofi/words/"

// Complete a selection
func Complete(selection string) {
	usr, _ := user.Current()
	home := usr.HomeDir
	words, err := wordsFromDir(home + wordsPath)
	if err != nil {
		fmt.Printf(err.Error())
	}
	completion, err := choose(selection, words)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf(completion)
}
