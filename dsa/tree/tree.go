package tree

import (
	"context"
	"fmt"
)

type node struct {
	value interface{}
	left  *node
	right *node
}
type tree struct {
	root *node
}

type Ops interface {
	Populate(ctx context.Context)
	Display(ctx context.Context)
	PreetyDisplay(ctx context.Context)
}

func NewTree() Ops {
	return &tree{}
}

func newNode(value interface{}) node {
	return node{
		value: value,
	}
}

func (t *tree) Populate(ctx context.Context) {
	if t.root == nil {
		fmt.Println("Enter the value of root node")
		var value int
		fmt.Scanln(&value)
		newNode := newNode(value)
		t.root = &newNode
		t.populate(ctx, t.root)
	}

}

func (t *tree) populate(ctx context.Context, node *node) {
	fmt.Printf("do you want to insert left of %v \n", node.value)
	var left bool
	fmt.Scanln(&left)
	if left {
		fmt.Println("Enter the value of left node")
		var value int
		fmt.Scanln(&value)
		newNode := newNode(value)
		node.left = &newNode
		t.populate(ctx, node.left)
	}

	fmt.Printf("do you want to insert right of %v \n", node.value)
	var right bool
	fmt.Scanln(&right)
	if right {
		fmt.Println("Enter the value of right node")
		var value int
		fmt.Scanln(&value)
		newNode := newNode(value)
		node.right = &newNode
		t.populate(ctx, node.right)
	}
}

func (t *tree) Display(ctx context.Context) {
	t.display(t.root, "")
}

func (t *tree) display(node *node, indend string) {
	if node == nil {
		return
	}
	fmt.Printf("%v %v\n", indend, node.value)

	t.display(node.left, indend+"\t")
	t.display(node.right, indend+"\t")
}

func (t *tree) PreetyDisplay(ctx context.Context) {
	t.preetyDisplay(ctx, t.root, 0)
}

func (t *tree) preetyDisplay(ctx context.Context, node *node, level int) {
	if node == nil {
		return
	}
	t.preetyDisplay(ctx, node.right, level+1)

	if level != 0 {
		for i := 0; i < level-1; i++ {
			fmt.Print("|\t\t")
		}
		fmt.Println("|-------------->", node.value)
	} else {
		fmt.Println(node.value)
	}

	t.preetyDisplay(ctx, node.left, level+1)

}
