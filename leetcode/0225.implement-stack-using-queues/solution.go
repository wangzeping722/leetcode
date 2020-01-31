package problem0225

type MyStack struct {
	datas []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		datas: []int{},
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.datas = append([]int{x}, this.datas...)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	data := this.datas[0]
	this.datas = this.datas[1:]
	return data
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.datas[0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.datas) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
