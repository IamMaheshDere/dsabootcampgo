package main

import (
	"context"

	"github.com/IamMaheshDere/dsabootcampgo/dsa/tree"
)

func main() {
	treeOps := tree.NewTree()
	treeOps.Populate(context.Background())
	treeOps.Display(context.Background())
	treeOps.PreetyDisplay(context.Background())
}
