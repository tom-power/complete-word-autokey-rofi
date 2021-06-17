package main

import (
	"./completeword"
	"flag"
	"os"
)

var completion = flag.String("complete", "", "a string")
var add = flag.String("add", "", "a string")

func main() {
	flag.Parse()
	switch os.Args[1] {
	case "complete":
		completeword.Complete(*completion)
	case "count":
		completeword.Add(*completion)
	}
}
