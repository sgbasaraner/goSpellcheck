package main

import "os"
import "fmt"
import "bufio"

func checkError (e error) {
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

	// decide which dictionary to load
	var dictionary string

	if len(args) != 3 {
		dictionary = "dictionaries/large"
	} else {
		dictionary = string(args[1])
	}

	// load the dictionary
	var loaded bool = load(dictionary)
	if !loaded {
		fmt.Println("Unable to load", dictionary)
		os.Exit(1)
	}

	// decide which text file to spellcheck
	var textFile string

	if len(args) != 3 {
		textFile = string(args[1])
	} else {
		textFile = string(args[2])
	}

	fmt.Println("\nMISSPELLED WORDS\n")

	// prepare to spellcheck
	words, misspellings := 0, 0
	text, err := os.Open(textFile)
	checkError(err)
	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		line := scanner.Text()
		words += 1
		if !check(line){
			fmt.Println(line)
			misspellings += 1
		}
	}

	fmt.Println("\nWORDS MISSPELLED:", misspellings)
	fmt.Println("WORDS IN DICTONARY:", size())
	fmt.Println("WORDS IN TEXT:", words)
	

}