package node

type Node struct {
	Data     interface{}
	Next     *Node
	Previous *Node
}
