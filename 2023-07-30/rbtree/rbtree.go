package rbtree

type NodeColor int

const (
	RedColor NodeColor = iota
	BlackColor
)

type Node struct {
	left, right *Node
	parent      *Node
	color       NodeColor
	Key         int
	value       interface{}
}

type RBTree struct {
	root *Node
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (rb *RBTree) Search(searchKey int) interface{} {
	cur := rb.root
	for cur != nil {
		if searchKey < cur.Key {
			cur = cur.left
			continue
		}
		if searchKey > cur.Key {
			cur = cur.right
			continue
		}
		return cur.value
	}
	return nil
}

func (rb *RBTree) Insert() {
	//todo
	panic("implement me ")
}
