package main

// TreeNode TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	allPossibleFBT(7)
}

func allPossibleFBT(n int) []*TreeNode {
	// total := n/3 + 1
	// // var possibleArr []int
	// half := TreeNode{
	// 	Val: 0,
	// }
	full := TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	var input TreeNode
	number := 3
	if input.Left.Left == nil || input.Right.Right == nil {
		input.Left = &full
		number += 2
		if number < n {
			input.Right = &full
		}
	}
	return nil
}
