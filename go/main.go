package main

import (
	"./completeword"
	"flag"
	"fmt"
	"os"
)

var wordToComplete = flag.String("complete", "", "a string")
var wordToAdd = flag.String("add", "", "a string")

func main() {
	flag.Parse()
	switch os.Args[1] {
	case "--complete":
		completion, err := completeword.Complete(
			completeword.WordsFromHomeDir,
			completeword.ChooseWordUsingRofi,
			*wordToComplete)
		if err != nil {
			fmt.Printf(*wordToComplete)
		}
		fmt.Printf(completion)
	case "--add":
		err := completeword.Add(*wordToAdd)
		if err != nil {
			fmt.Printf(err.Error())
		}
	default:
		fmt.Printf("usage --complete <partialWord> | --add <partialWord>")
	}
}
