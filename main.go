package main

import "os"
import "fmt"
import "io/ioutil"

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

	dat, err := ioutil.ReadFile(args[1])
	check(err)
	fmt.Print(string(dat))
}