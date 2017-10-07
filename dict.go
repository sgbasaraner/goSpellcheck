package main
import "fmt"

func check(word string) bool {
	return true	
}

func load(dictionary string) bool {
	root := &(trieNode{label: 0})
	root.AddWord("araba")
	if root.Search("araba") {
		fmt.Println("araba exists")
	}
	// if root.Search("arap") {
	// 	fmt.Println("arap exists")
	// }
	return true
}

func size() int {
	return 0
}

