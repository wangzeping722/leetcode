package _447_number_of_boomerangs

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		m := make(map[int]int, len(points))
		for j := 0; j < len(points); j++ {
			if i != j {
				dis := dist(points[i], points[j])
				res += m[dis]*2
				m[dis]++
			}
		}
	}
	return res
}

func dist(i, j []int) int {
	return (i[0]-j[0])*(i[0]-j[0]) + (i[1]-j[1])*(i[1]-j[1])
}
