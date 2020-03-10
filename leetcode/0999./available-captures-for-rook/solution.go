package available_captures_for_rook

import "strings"

func numRookCaptures(board [][]byte) int {
	ri, rj := -1, -1
	for i := 0; i < len(board); i++ {
		if ri > -1 {
			break
		}
		for j := 0; j < len(board[0]); j++{
			if board[i][j] == 'R' {
				ri, rj = i, j
			}
		}
	}
	if ri == -1 {
		return 0
	}

	res := 0
	// 水平方向
	hStr := strings.ReplaceAll(string(board[ri]), ".", "")
	if strings.Contains(hStr, "pR") {
		res++
	}
	if strings.Contains(hStr, "Rp") {
		res++
	}

	vStr := ""
	for i:=0; i< len(board); i++ {
		vStr += string(board[i][rj])
	}
	vStr = strings.ReplaceAll(vStr, ".", "")
	if strings.Contains(vStr, "pR") {
		res++
	}
	if strings.Contains(vStr, "Rp") {
		res++
	}
	return res
}
