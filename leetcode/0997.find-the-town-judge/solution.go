package _997_find_the_town_judge

func findJudge(N int, trust [][]int) int {
	inDegree := make([]int, N+1)
	outDegree := make([]int, N+1)
	for _, path := range trust {
		outDegree[path[0]]++
		inDegree[path[1]]++
	}

	for i:=1; i<=N; i++ {
		if inDegree[i] == N-1 && outDegree[i] == 0 {
			return i
		}
	}

	return -1
}
