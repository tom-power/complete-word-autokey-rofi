package completeword

import (
	"bufio"
	"io"
	"os"
	"strings"
  "io/ioutil"
	"fmt"
)

func words(dir string) (string, error) {
	files, err := ioutil.ReadDir(dir); if err != nil { return "", err }
	return wordsInFiles(files)
}

func wordsInFiles(files []os.FileInfo) (string, error) {
	var words string
	for _, file := range files {
		fmt.Printf(file.Name())
		fileWords, err := wordsInFile(file.Name()); if err != nil { return "", err }
		words = words + fileWords
	}
	return words, nil
}

func wordsInFile(file string) (string, error) {
	handle, err := os.Open(file);  if err != nil { return "", err }
	defer handle.Close()
  return readWords(handle)
}

func readWords(handle io.Reader) (string, error) {
	var words []string
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}	
	return strings.Join(words, " "), scanner.Err()
}
