package _303_range_sum_query_immutable

type NumArray struct {
	datas []int
}


func Constructor(nums []int) NumArray {
	d := make([]int, len(nums))
	if len(nums) == 0 {
		return NumArray{datas:d}
	}
	d[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		d[i] = d[i-1] + nums[i]
	}
	return NumArray{datas:d}
}


func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.datas[j]
	}
	return this.datas[j] - this.datas[i-1]
}

/*
type NumArray struct {
	Data []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return NumArray{nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.Data[j]
	} else {
		return this.Data[j] - this.Data[i-1]
	}
}
 */

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
