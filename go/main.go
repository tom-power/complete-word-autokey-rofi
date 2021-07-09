package main

import (
	"./completeword"
	"flag"
	"fmt"
	"os"
)

var wordToComplete = flag.String("complete", "", "a string")
var wordToAdd = flag.String("add", "", "a string")

var complete = completeword.Complete(
  completeword.WordsFromHomeDir,
  completeword.ChooseWordUsingRofi,
)

var add = completeword.Add(
  completeword.WordsFromHomeDir,
  completeword.ChooseWordUsingRofi,
  completeword.AddedPath,
  completeword.WriteWordToPath,
)

func main() {
	flag.Parse()
	switch os.Args[1] {
	case "--complete":
		completion, err := complete(*wordToComplete)
		if err != nil {
			fmt.Printf(*wordToComplete)
		}
		fmt.Printf(completion)
	case "--add":
		err := add(*wordToAdd)
		if err != nil {
			fmt.Printf(err.Error())
		}
	default:
		fmt.Printf("usage --complete <partialWord> | --add <partialWord>")
	}
}
