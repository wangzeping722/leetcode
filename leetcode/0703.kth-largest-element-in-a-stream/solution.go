package main

type KthLargest struct {
	pq      []int
	size    int
	maxSize int
}

func Constructor(k int, nums []int) KthLargest {
	heap := KthLargest{
		pq:      make([]int, k+1),
		size:    0,
		maxSize: k,
	}

	for _, num := range nums {
		heap.Add(num)
	}

	return heap
}

func (this *KthLargest) Add(val int) int {
	if this.size < this.maxSize {
		this.size++
		this.pq[this.size] = val
		this.Swim(this.size)
	} else {
		if this.pq[1] >= val {
			return this.pq[1]
		} else {
			this.pq[1] = val
			this.Sink(1) //下沉
		}
	}
	return this.pq[1]
}

func (this *KthLargest) Sink(k int) {
	for 2*k <= this.size {
		// 选择左右节点中小的节点，替换到当前节点
		older := this.Left(k)
		if this.Right(k) <= this.maxSize && this.Less(this.Right(k), older) {
			older = this.Right(k)
		}
		if this.Less(k, older) {
			break
		}
		this.pq[k], this.pq[older] = this.pq[older], this.pq[k]
		k = older
	}
}

// 上浮第k个元素，以维护最小堆的性质
func (this *KthLargest) Swim(k int) {
	for k > 1 && this.Less(k, this.Parent(k)) {
		this.pq[k], this.pq[this.Parent(k)] = this.pq[this.Parent(k)], this.pq[k]
		k = this.Parent(k)
	}
}

func (this *KthLargest) Less(i, j int) bool {
	return this.pq[i] < this.pq[j]
}

// 父节点的索引
func (this *KthLargest) Parent(root int) int {
	return root / 2
}

// 左孩子索引
func (this *KthLargest) Left(root int) int {
	return root * 2
}

// 右孩子索引
func (this *KthLargest) Right(root int) int {
	return root*2 + 1
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
