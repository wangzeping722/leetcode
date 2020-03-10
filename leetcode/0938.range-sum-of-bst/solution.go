package _938_range_sum_of_bst

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Val >= L && root.Val<= R {
			sum += root.Val
		}

		if root.Val < L {
			dfs(root.Right)
			return
		}
		if root.Val > R {
			dfs(root.Left)
			return
		}
		dfs(root.Right)
		dfs(root.Left)
	}
	dfs(root)
	return sum
}

func rangeSumBST1(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST1(root, L, R)
	}
	if root.Val > R {
		return rangeSumBST1(root, L, R)
	}
	return root.Val + rangeSumBST1(root.Left, L, R) + rangeSumBST1(root.Right, L, R)
}