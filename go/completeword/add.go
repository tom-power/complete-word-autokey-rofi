package completeword

import (
	"os"
	"os/user"
)

const addedPath = "/.config/complete-word-autokey-rofi/words/added.txt"

func Add(selection string) error {
	usr, _ := user.Current()
	addPath := usr.HomeDir + addedPath
	words, err := wordsFromFile(addPath)
	if err != nil {
		return err
	}
	word, err := ChooseWordUsingRofi(words, selection)
	if err != nil {
		return err
	}
	return writeWord(addPath, word)
}

func writeWord(path string, word string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(word))
	return err

}
