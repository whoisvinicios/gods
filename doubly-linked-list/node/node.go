package node

type Node struct {
	Data     interface{}
	Next     *Node
	Previous *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data:     data,
		Next:     nil,
		Previous: nil,
	}
}
