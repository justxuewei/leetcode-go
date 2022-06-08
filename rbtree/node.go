package rbtree

type color int

const (
	red color = iota
	black
)

type Node struct {
	Key                 int
	color               color
	Left, Right, Parent *Node
}
