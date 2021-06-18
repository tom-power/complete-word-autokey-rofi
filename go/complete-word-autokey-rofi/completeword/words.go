package completeword

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"os/user"
)

type GetWords func() ([]string, error)

const wordsConfigPath = "/.config/complete-word-autokey-rofi/words/"

var WordsFromHomeDir = func() ([]string, error) { usr, _ := user.Current()
	return wordsFromDir(usr.HomeDir + wordsConfigPath)
}

func wordsFromDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}
	return wordsFromFiles(filePaths(dir, files))
}

func filePaths(dir string, files []os.FileInfo) []string {
	var filePaths []string
	for _, file := range files {
		filePaths = append(filePaths, dir+file.Name())
	}
	return filePaths
}

func wordsFromFiles(filePaths []string) ([]string, error) {
	var words []string
	for _, filePath := range filePaths {
		fileWords, err := wordsFromFile(filePath)
		if err != nil {
			return words, err
		}
	  words = append(words, fileWords...)
	}
	return words, nil
}

func wordsFromFile(filePath string) ([]string, error) {
	fileReader, err := os.Open(filePath)
	if err != nil {
	  return []string{}, err
	}
	fileWords, err := readWords(fileReader)
	if err != nil {
		return []string{}, err
	}
	return fileWords, err
}

func readWords(handle io.Reader) ([]string, error) {
	var words []string
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

func writeWord(path string, word string) error {
	handle, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := handle.Write([]byte(word)); err != nil {
		return err
	}
	return nil
}
