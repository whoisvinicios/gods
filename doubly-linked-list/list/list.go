package list

import (
	"fmt"
	"gods/doubly-linked-list"
)

type List struct {
	Head   *node.Node
	Tail   *node.Node
	Length int
}

func CreateList() *List {
	return &List{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (l *List) Add(n *node.Node) error {
	if l.Head == nil {
		l.Head = n
	} else {
		currentNote := l.Tail
		currentNote.Next = n
		n.Previous = l.Tail
	}
	l.Tail = n
	return nil
}

func (l *List) PrintAllNodes() error {
	if l.Head == nil {
		return fmt.Errorf("list is empty")
	}
	currentNode := l.Head
	for currentNode.Next != nil {
		fmt.Printf("%+v", *currentNode)
		currentNode = currentNode.Next
	}
	return nil
}
