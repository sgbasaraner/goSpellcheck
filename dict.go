package main

func check(word string) bool {
	return true	
}

func load(dictionary string) bool {
	root := &(trieNode{label: 0})
	root.AddChild(byte('a'))
	root.PrintChildren()
	return true
}

func size() int {
	return 0
}

