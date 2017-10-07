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