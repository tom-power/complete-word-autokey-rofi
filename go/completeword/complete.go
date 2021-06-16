package completeword

import (
	"fmt"
	"os/user"
)

// Complete is ...
func Complete(selection string) {
	usr, _ := user.Current()
	dir := usr.HomeDir
	words, err := words(dir + "/.config/complete-words-autokey-rofi/words/")
	completion, err := choose(selection, words)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf(completion)
}
