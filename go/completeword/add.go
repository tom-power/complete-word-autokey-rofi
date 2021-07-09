package completeword

import (
	"errors"
	"os"
	"os/user"
)

const addedPath = "/.config/complete-word-autokey-rofi/words/added.txt"

type GetPath = func() (string, error)

var AddedPath GetPath = func() (string, error) {
	usr, error := user.Current()
	if error != nil {
		return "", error
	}
	return usr.HomeDir + addedPath, nil
}

type WordsFrom = func(string) ([]string, error)

var WordsFromPath = func(path string) ([]string, error) {
	return wordsFromFile(path)
}

type WriteWord func(string, string) error

var WriteWordToPath = func(path string, word string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(word))
	return err
}

func Add(
	getWords GetWords,
	chooseWord ChooseWord,
	addPath GetPath,
	writeWord WriteWord,
) func(selection string) error {
  return func(selection string ) error {
    words, err := getWords()
    if err != nil {
      return err
    }
    word, err := chooseWord(words, selection)
    if err != nil {
      return err
    }
    if isIn(word, words) {
      return errors.New("can't add duplicate word")
    }
    path, err := addPath()
    if err != nil {
      return err
    }
    return writeWord(path, word)
  }
}

func isIn(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
