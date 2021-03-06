package _095

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return dfs(1, n)
}

func dfs(start, end int) []*TreeNode {
	var res []*TreeNode
	if start > end {
		res = append(res, nil)
		return res
	}

	for i:= start; i <= end; i++ {
		for _, l := range dfs(start, i - 1) {
			for _, r := range dfs(i+1, end) {
				cur := &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				}
				res = append(res, cur)
			}
		}
	}

	return res
}
