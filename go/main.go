package main

import (
	"flag"
	"./completeword"
)

var completion = flag.String("complete", "", "a string")

func main() {
	flag.Parse()
	completeword.Complete(*completion)
}