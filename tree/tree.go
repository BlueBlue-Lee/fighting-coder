package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%v ", root.Val)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.Left)
	PostOrderTraversal(root.Right)
	fmt.Printf("%v ", root.Val)
}

func InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left)
	fmt.Printf("%v ", root.Val)
	InOrderTraversal(root.Right)
}

func LevelOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		first := nodes[0]
		fmt.Printf("%v ", first.Val)
		if first.Left != nil {
			nodes = append(nodes, first.Left)
		}
		if first.Right != nil {
			nodes = append(nodes, first.Right)
		}
		nodes = nodes[1:]
	}
}

func LevelOrderBottomTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	levels := make([][]*TreeNode, 0)
	levels = append(levels, []*TreeNode{root})
	for {
		lastLevel := levels[len(levels)-1]
		var newLevel []*TreeNode
		for _, n := range lastLevel {
			if n.Left != nil {
				newLevel = append(newLevel, n.Left)
			}
			if n.Right != nil {
				newLevel = append(newLevel, n.Right)
			}
		}
		if len(newLevel) == 0 {
			break
		}
		levels = append(levels, newLevel)
	}
	for i := len(levels) - 1; i >= 0; i-- {
		for _, n := range levels[i] {
			fmt.Printf("%v ", n.Val)
		}
	}
}
func GetKLevelNodes(root *TreeNode, k int) []int {
	var result []int
	if root == nil {
		return result
	}
	depth := 1
	if k == 1 {
		return []int{root.Val}
	}
	levels := make([][]*TreeNode, 0)
	levels = append(levels, []*TreeNode{root})
	for {
		depth++
		lastLevel := levels[len(levels)-1]
		newLevels := make([]*TreeNode, 0)
		for _, n := range lastLevel {
			if n.Left != nil {
				newLevels = append(newLevels, n.Left)
			}
			if n.Right != nil {
				newLevels = append(newLevels, n.Right)
			}
		}
		if len(newLevels) == 0 {
			break
		}
		levels = append(levels, newLevels)
		if depth == k {
			for _, n := range newLevels {
				result = append(result, n.Val)
			}
			return result
		}
	}
	return result
}

// GetTreeDepth return depth of tree
func GetTreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1 + GetTreeDepth(root.Left)
	right := 1 + GetTreeDepth(root.Right)
	if left > right {
		return left
	}
	return right
}

var longest int

func GetLongestPath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	getHeight(root)
	return longest
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getHeight(root.Left)
	right := getHeight(root.Right)
	if left+right > longest {
		longest = left + right
	}
	bigger := left
	if right > bigger {
		bigger = right
	}
	return bigger + 1
}

func GetMaxSum(root *TreeNode) int {
	return 0
}

func ZigzagLevelTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	levels := make([][]*TreeNode, 0)
	levels = append(levels, []*TreeNode{root})
	for {
		lastLevel := levels[len(levels)-1]
		var newLevel []*TreeNode
		for _, n := range lastLevel {
			if n.Left != nil {
				newLevel = append(newLevel, n.Left)
			}
			if n.Right != nil {
				newLevel = append(newLevel, n.Right)
			}
		}
		if len(newLevel) == 0 {
			break
		}
		levels = append(levels, newLevel)
	}
	for i := range levels {
		if i%2 == 0 {
			for _, n := range levels[i] {
				result = append(result, n.Val)
			}
		} else {
			for j := len(levels[i]) - 1; j >= 0; j-- {
				result = append(result, levels[i][j].Val)
			}
		}
	}
	return result
}

func (t *TreeNode) BinarySearch(target int) {

}

func (t *TreeNode) MaxDepth() int {
	return 0
}

func (t *TreeNode) MinDepth() int {
	return 0
}

func (t *TreeNode) GetNodesCount() int {
	return 0
}

func (t *TreeNode) GetKLevelNodesCount(k int) int {
	return 0
}

func (t *TreeNode) GetKLevelLeafNodesCount(k int) int {
	return 0
}

func (t *TreeNode) IsSameTree(s *TreeNode) bool {
	return false
}

func (t *TreeNode) IsBalanced() bool {
	return false
}

func (t *TreeNode) GetMirrorTree() *TreeNode {
	return nil
}

func (t *TreeNode) GetLowestCommonAncestor(a *TreeNode, b *TreeNode) *TreeNode {
	return nil
}
