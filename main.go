package main

import (
	"context"

	"github.com/IamMaheshDere/dsabootcampgo/dsa/tree"
)

func main() {
	tree := tree.NewTree()
	tree.Populate(context.Background())
	tree.Display(context.Background())
	tree.PreetyDisplay(context.Background())
}
