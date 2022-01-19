package main

import (
	"fmt"
)

type Node struct {
	Previous *Node
	Data     int
	Next     *Node
}

var tail *Node

func (n *Node) AddNode(data int) {
	fmt.Println("Inserting node", data)
	if tail == nil {
		tail = n
	}
	n = tail
	newNode := Node{n, data, nil}
	n.Next = &newNode
	tail = &newNode
}

func (n *Node) PrintNode() {
	fmt.Println("Printing nodes")
	iter := n
	for iter != nil {
		fmt.Println(iter.Data)
		iter = iter.Next
	}
}

func (n *Node) DeleteLast() {
	fmt.Println("Deleting last node", tail.Data)
	tail = tail.Previous
	(tail.Next).Previous = nil
	tail.Next = nil
}
