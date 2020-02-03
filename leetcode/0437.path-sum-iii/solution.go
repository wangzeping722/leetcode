package _437_path_sum_iii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return paths(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func paths(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	res := 0
	sum -= root.Val
	if sum == 0 {
		res++
	}
	return res + paths(root.Left, sum) + paths(root.Right, sum)
}

// 很明显的优化，使用数组缓存
func pathSum1(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	paths := make([]int, 1000)
	return pathsum(root, sum, paths, 0)
}

func pathsum(root *TreeNode, sum int, paths []int, p int) int {
	if root == nil {
		return 0
	}

	paths[p] = root.Val

	res, tmp := 0, 0
	for i := p; i >= 0; i-- {
		tmp += paths[i]
		if tmp == sum {
			res++
		}
	}
	return res + pathsum(root.Left, sum, paths, p+1) + pathsum(root.Right, sum, paths, p+1)
}
