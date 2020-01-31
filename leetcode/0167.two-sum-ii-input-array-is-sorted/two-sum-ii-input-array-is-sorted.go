package problem0167

func twoSum(numbers []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(numbers); i++ {
		if n, ok := m[target-numbers[i]]; ok {
			return []int{n, i + 1}
		}
		m[numbers[i]] = i + 1
	}
	return nil
}

func twoSum1(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else if sum > target {
			j--
		}
	}
	return nil
}
