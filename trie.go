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
	for i := 0; i < len(n.children); i++ {
		fmt.Println(string(n.children[i].label))
	}
}

func (n *trieNode) HasChild(l byte) *trieNode {
	for i := 0; i < len(n.children); i++ {
		if l == n.children[i].label {
			return (n.children[i])
		}
	}
	return nil
}

func (n *trieNode) AddWord(word string) {
	cursor := n
	for i := 0; i < len(word); i++ {
		if cursor.HasChild(byte(word[i])) == nil {
			// fmt.Println("cursor is:", cursor.label)
			cursor = cursor.AddChild(byte(word[i]))
		} else {
			cursor = cursor.HasChild(byte(word[i]))
		}
	}
	// tilde indicates end of the word
	cursor.AddChild(byte('~'))
}

func (n *trieNode) Search(word string) bool {
	cursor := n
	for i := 0; i < len(word); i++ {
		// cursor.PrintChildren()
		if cursor.HasChild(byte(word[i])) != nil {
			cursor = cursor.HasChild(byte(word[i]))
		} else {
			return false
		}
	}

	if cursor.HasChild(byte('~')) != nil {
		return true
	} else {
		return false
	}
}