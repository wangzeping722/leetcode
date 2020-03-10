package interview

type MaxQueue struct {
	maxQueue []int
	datas    []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		maxQueue: []int{},
		datas:    []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.datas) == 0 {
		return -1
	}
	return this.maxQueue[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.datas = append(this.datas, value)
	// 如果为 0
	if len(this.datas) == 0 {
		this.maxQueue = append(this.maxQueue, value)
	} else {
		if value >= this.maxQueue[0] {
			this.maxQueue = []int{}
			this.maxQueue = append(this.maxQueue, value)
			return
		}
		// 否则找到别当前大的值的后一位, 切掉后面的值,
		// 这样可以使最大队列永远是一个递减的状态
		for i := len(this.maxQueue) - 1; this.maxQueue[i] < value; i-- {
			this.maxQueue = this.maxQueue[:i]
		}

		this.maxQueue = append(this.maxQueue, value)
	}
}

func (this *MaxQueue) Pop_front() int {
	if len(this.datas) == 0 {
		return -1
	}
	ret := this.datas[0]
	if ret == this.maxQueue[0] {
		this.maxQueue = this.maxQueue[1:]
	}
	this.datas = this.datas[1:]
	return ret
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
