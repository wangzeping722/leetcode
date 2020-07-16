package main

func isBipartite(graph [][]int) bool {
	// 标记所有节点是否访问过
	visited := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		// 如果访问过, 那么就跳过这个节点
		if visited[i] != 0 {
			continue
		}
		// 访问队列
		queue := make([]int, 0)
		queue = append(queue, i)
		visited[i] = 1 // 染色为蓝色
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			curColor := visited[cur]
			neighborColor := -curColor
			for j := 0; j < len(graph[cur]); j++ {
				neighbor := graph[cur][j]
				if visited[neighbor] == 0 {
					visited[neighbor] = neighborColor
					queue = append(queue, neighbor)
				} else if visited[neighbor] != neighborColor {
					return false
				}
			}
		}
	}
	return true
}

func main() {

}
