package problem0232

type MyQueue struct {
	datas []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{datas: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.datas = append(this.datas, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	data := this.datas[0]
	this.datas = this.datas[1:]
	return data
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.datas[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.datas) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
