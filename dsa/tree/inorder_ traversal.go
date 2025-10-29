package tree

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/

func inorderTraversal(root *TreeNode) []int {
	numbers := make([]int, 0)

	if root != nil {
		numbers = inorderTraversalRec(root, numbers)
	}

	return numbers

}

func inorderTraversalRec(node *TreeNode, numbers []int) []int {
	if node.Left != nil {
		numbers = inorderTraversalRec(node.Left, numbers)
	}

	numbers = append(numbers, node.Val)

	if node.Right != nil {
		numbers = inorderTraversalRec(node.Right, numbers)
	}

	return numbers
}
