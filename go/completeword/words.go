package completeword

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func words(dir string) (string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}
	return wordsInFiles(filePaths(dir, files))
}

func filePaths(dir string, files []os.FileInfo) []string {
	var filePaths []string
	for _, file := range files {
		filePaths = append(filePaths, dir+file.Name())
	}
	return filePaths
}

func wordsInFiles(filePaths []string) (string, error) {
	var words string
	for _, filePath := range filePaths {
		fileReader, err := os.Open(filePath)
		if err != nil {
			return "", err
		}
		fileWords, err := readWords(fileReader)
		if err != nil {
			return "", err
		}
		words = words + fileWords
	}
	return words, nil
}

func readWords(handle io.Reader) (string, error) {
	var words []string
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return strings.Join(words, "\n"), scanner.Err()
}
