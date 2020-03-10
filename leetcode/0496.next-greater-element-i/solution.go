package _496_next_greater_element_i

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	indexOf := make(map[int]int)
	for i, n := range nums2 {
		indexOf[n] = i
	}

	res := make([]int, len(nums1))
	for i, n := range nums1 {
		res[i] = -1
		for j := indexOf[n] + 1; j < len(nums2); j++ {
			if n < nums2[j] {
				res[i] = nums2[j]
				break
			}
		}
	}
	return res
}

type stack struct {
	s []int
}

func (s *stack) Pop() (int, error) {
	if len(s.s) > 0 {
		v := s.s[len(s.s) - 1]
		s.s = s.s[0:len(s.s) - 1]
		return v, nil
	} else {
		return 0, fmt.Errorf("null")
	}
}

func (s *stack) Push(v int) {
	s.s = append(s.s, v)
}

func (s *stack) Top() (int, error) {
	if len(s.s) > 0 {
		return s.s[len(s.s) - 1], nil
	} else {
		return 0, fmt.Errorf("null")
	}
}

func nextGreaterElementStack(nums1 []int, nums2 []int) []int {
	var s stack
	m := make(map[int]int, len(nums2))
	for _, v := range nums2 {
		for {
			mv, err := s.Top()
			if (err != nil) || (err == nil && mv >= v) {
				break
			}
			if err == nil && v > mv{
				mv, err = s.Pop()
				m[mv] = v
			}
		}
		s.Push(v)
	}
	for {
		mv, err := s.Pop()
		if err != nil {
			break
		} else {
			m[mv] = -1
		}
	}
	var re []int
	for _, v := range nums1 {
		if index, ok := m[v]; ok {
			re = append(re, index)
		}
	}
	return re
}