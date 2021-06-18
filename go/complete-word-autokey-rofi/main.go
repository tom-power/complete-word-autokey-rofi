package main

import (
	"complete-word-autokey-rofi/completeword"
	"flag"
	"os"
)

var wordToComplete = flag.String("complete", "", "a string")
var wordToAdd = flag.String("add", "", "a string")

func main() {
	flag.Parse()
	switch os.Args[1] {
  	case "complete":
  		completeword.Complete(
  		  completeword.WordsFromDir,
  		  completeword.ChooseWordUsingRofi,
  		  *wordToComplete)
  	case "add":
  		completeword.Add(*wordToAdd)
	}
}
