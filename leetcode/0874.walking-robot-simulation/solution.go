package _874_walking_robot_simulation

var DIRECT_POINTS = []*point{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func robotSim(commands []int, obstacles [][]int) int {
	var bad = map[int]bool{}
	var bad2 = map[int]map[int]bool{}
	// 遍历obstacles
	for _, obstacle := range obstacles {
		bad[obstacle[0]] = true
		if _, exsits := bad2[obstacle[0]]; !exsits {
			bad2[obstacle[0]] = map[int]bool{}
		}
		bad2[obstacle[0]][obstacle[1]] = true
	}
	max := 0
	p := point{0, 0}
	directIdx := 0
	for _, command := range commands {
		switch command {
		case -2:
			directIdx = (directIdx - 1 + 4) % 4
		case -1:
			directIdx = (directIdx + 1) % 4
		default:
			for i:=0; i<command; i++ {
				p.forward(DIRECT_POINTS[directIdx])
				if bad[p.X] && bad2[p.X][p.Y] {
					p.backward(DIRECT_POINTS[directIdx])
					break
				}
				if p.length() > max {
					max = p.length()
				}
			}
		}
	}
	return max
}

type point struct {
	X int
	Y int
}

func (p *point) forward(p2 *point) {
	p.X += p2.X
	p.Y += p2.Y
}

func (p *point) backward(p2 *point) {
	p.X -= p2.X
	p.Y -= p2.Y
}

func (p *point) length() int {
	return p.X*p.X + p.Y*p.Y
}
