package main

import "os"
import "fmt"

func check (e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// check command line arguments' count
	args := os.Args
	if len(args) != 2 && len(args) != 3 {
		fmt.Println("Usage: spellcheck [dictionary] text")
	}

	var dictionary string

	if len(args) != 3 {
		dictionary = "dictionaries/large"
	} else {
		dictionary = string(args[2])
	}
}