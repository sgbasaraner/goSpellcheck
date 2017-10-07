package main

import "fmt"

type trieNode struct {
	label 		byte
	children 	[]*trieNode
}

func (n *trieNode) AddChild(l byte) *trieNode {
	node := trieNode{l, nil}
	np := &node
	n.children = append(n.children, np)
	return np
}

func (n *trieNode) PrintChildren() {
	fmt.Println(len(n.children))
	for i := 0; i < len(n.children); i++ {
		fmt.Printf(string(n.children[i].label))
	}
}

func (n *trieNode) HasChild(l byte) bool {
	for i := 0; i < len(n.children); i++ {
		if l == n.children[i].label {
			return true
		}
	}
	return false
}

func (n *trieNode) AddWord(word string) {
	cursor := n
	for i := 0; i < len(word); i++ {
		if !(cursor.HasChild(byte(word[i]))) {
			fmt.Println(word[i])
			cursor = cursor.AddChild(byte(word[i]))
		}
	}
	// tilde indicates end of the word
	cursor.AddChild(byte('~'))
}