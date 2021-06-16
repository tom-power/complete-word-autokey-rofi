package completeword

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func words(file string) string {
	handle, error := os.Open(file)
	if error != nil {
		fmt.Printf("Failed to read file: %s", file)
	}
	defer handle.Close()
  words, error := readWords(handle)
  if error != nil {
		fmt.Printf("Failed to read words: %s", file)
	}
	return words
}

func readWords(handle io.Reader) (string, error) {
	var words []string
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return strings.Join(words, " "), scanner.Err()
}
