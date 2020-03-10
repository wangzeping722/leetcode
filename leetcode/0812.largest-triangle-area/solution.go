package _812_largest_triangle_area

import "math"

func largestTriangleArea(points [][]int) float64 {
	var calc func(a, b, c []int) float64
	calc = func(a, b, c []int) float64 {
		return math.Abs(0.5 * float64(a[0]*b[1]+b[0]*c[1]+c[0]*a[1]-a[1]*b[0]-b[1]*c[0]-c[1]*a[0]))
	}
	res := 0.0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				temp := calc(points[i], points[j], points[k])
				if temp > res {
					res = temp
				}
			}
		}
	}
	return res
}

func largestTriangleArea1(points [][]int) float64 {
	var calc func(a, b, c []int) float64
	calc = func(a, b, c []int) float64 {
		return math.Abs(0.5 * float64(a[0]*b[1]+b[0]*c[1]+c[0]*a[1]-a[1]*b[0]-b[1]*c[0]-c[1]*a[0]))
	}

	res := 0.0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				temp := calc(points[i], points[j], points[k])
				if temp > res {
					res = temp
				}
			}
		}
	}
	return res
}