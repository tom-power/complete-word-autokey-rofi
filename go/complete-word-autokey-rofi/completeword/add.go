package completeword

import (
	"fmt"
	"os/user"
)

const addPath = "/.config/complete-word-autokey-rofi/words/add.txt"

// Add a word
func Add(selection string) {
	usr, _ := user.Current()
	home := usr.HomeDir
	words, err := wordsFromFile(home + addPath)
	if err != nil {
		fmt.Printf(err.Error())
	}
	word, err := ChooseWordUsingRofi(selection, words)
	if err != nil {
		fmt.Printf(err.Error())
	}
	writeWord(addPath, word)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
