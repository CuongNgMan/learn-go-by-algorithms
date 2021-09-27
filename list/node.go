package list

type Node struct {
	data interface{}
	next *Node
}

type DNode struct {
	data interface{}
	next *DNode
	prev *DNode
}