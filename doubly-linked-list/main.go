package main

import (
	"gods/doubly-linked-list/list"
	"gods/doubly-linked-list/node"
	"log"
)

func main() {
	dlist := list.CreateList()

	node1 := node.NewNode(1)
	dlist.Add(node1)

	node2 := node.NewNode(2)
	dlist.Add(node2)

	node3 := node.NewNode(4)
	dlist.Add(node3)

	if err := dlist.PrintAllNodes(); err != nil {
		log.Fatalf("%s", err)
	}
}
