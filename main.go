package main

import "os"
import "fmt"
import "bufio"
import "regexp"
import "strings"

var lean bool = false

func checkError (e error) {
	if e != nil {
		panic(e)
	}
}

func isAlpha(str string) bool {
	for i := range str {
		if str[i] < 'A' || str[i] > 'z' {
			return false
		} else if str[i] > 'Z' && str[i] < 'a' {
			return false
		}
	}
	return true
}

func main() {

	// check command line arguments' count
	args := os.Args
	argLength := len(args)

	if argLength > 4 || argLength < 3 || (argLength == 4 && args[3] != "lean") {
		fmt.Println("Usage: spellcheck text dictionary [lean]")
		os.Exit(2)
	}

	// evaluate args
	// decide which dictionary & text will be loaded
	// set to lean mode (optional)
	var dictionary string = args[2]
	var textFile string = args[1]
	if argLength == 4{
		lean = true
	}

	// load the dictionary
	var loaded bool = load(dictionary)
	if !loaded {
		fmt.Println("Unable to load", dictionary)
		os.Exit(1)
	}

	fmt.Println("\nMISSPELLED WORDS\n")

	// prepare to spellcheck
	words, misspellings := 0, 0
	text, err := os.Open(textFile)
	checkError(err)
	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)

	if lean {
		for scanner.Scan() {
			line := scanner.Text()
			if !check(line){
				fmt.Println(line)
				misspellings += 1
			}
			words += 1
		}
	} else {
		for scanner.Scan() {
			line := scanner.Text()

			// ignore words with numbers like MS Word
			match, err := regexp.MatchString("[0-9]+", line)
			checkError(err)

			// reconstruct the word to take only alphabetic characters and apostrophe
			var word string
			for i := 0; i < len(line); i++ {
				if isAlpha(string(line[i])) || (line[i] == '\'') {
					word = word + string(line[i])
				}
			}

			word = strings.TrimSpace(word)
			// ignore too long words
			if !match && len(word) <= 45{
				if !check(word) && word != ""{
					fmt.Println(word)
					misspellings += 1
				}
			}
			words += 1
		}
	}
	

	fmt.Println("\nWORDS MISSPELLED:", misspellings)
	fmt.Println("WORDS IN DICTONARY:", size())
	fmt.Println("WORDS IN TEXT:", words)
	

}