package main
import "bufio"
import "os"
import "strings"

var words int = 0
var root *trieNode = &(trieNode{0, nil})

func check(word string) bool {
	return root.Search(strings.ToLower(word))	
}

func load(dictionary string) bool {
	text, err := os.Open(dictionary)
	checkError(err)
	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		line := scanner.Text()
		words += 1
		root.AddWord(line)
	}

	return true
}

func size() int {
	return words
}

