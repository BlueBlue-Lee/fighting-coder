package tree

import (
	"fmt"
	"testing"
)

/*
//                   10
//                /      \
//               20      130
//              /       /
//             30      140
//           /    \      \
//          40     22     190
//            \    /
//            90  50
*/
var root = &TreeNode{
	Val: 10,
	Left: &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val: 30,
			Left: &TreeNode{
				Val: 40,
				Right: &TreeNode{
					Val: 90,
				},
			},
			Right: &TreeNode{
				Val: 22,
				Left: &TreeNode{
					Val: 50,
				},
			},
		},
	},
	Right: &TreeNode{
		Val: 130,
		Left: &TreeNode{
			Val: 140,
			Right: &TreeNode{
				Val: 190,
			},
		},
	},
}

func TestGetLongestPath(t *testing.T) {
	fmt.Printf("%v", GetLongestPath(root))
}

func TestZigzagLevelTraversal(t *testing.T) {
	fmt.Printf("%v", ZigzagLevelTraversal(root))
}

func TestGetTreeDepth(t *testing.T) {
	fmt.Printf("%v", GetTreeDepth(root))
}

func TestGetKLevelNodes(t *testing.T) {
	fmt.Printf("%v\n", GetKLevelNodes(root, 1))
	fmt.Printf("%v\n", GetKLevelNodes(root, 2))
	fmt.Printf("%v\n", GetKLevelNodes(root, 3))
	fmt.Printf("%v\n", GetKLevelNodes(root, 5))
	fmt.Printf("%v\n", GetKLevelNodes(root, 8))
}

func TestPreOrderTraversal(t *testing.T) {
	// 10 20 30 40 90 22 50 130 140 190
	PreOrderTraversal(root)
}

func TestPostOrderTraversal(t *testing.T) {
	// 90 40 50 22 30 20 190 140 130 10
	PostOrderTraversal(root)
}

func TestInOrderTraversal(t *testing.T) {
	// 40 90 30 50 22 20 10 140 190 130
	InOrderTraversal(root)
}

func TestLevelOrderTraversal(t *testing.T) {
	// 10 20 130 30 140 40 22 190 90 50
	LevelOrderTraversal(root)
}

func TestLevelOrderBottomTraversal(t *testing.T) {
	// 90 50 40 22 190 30 140 20 130 10
	LevelOrderBottomTraversal(root)
}
