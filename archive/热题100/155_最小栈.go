package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 8:34 下午
 * @Desc:
 **/

type MinStack struct {
	dataStack []int
	minStack  []int
}

func Constructor() MinStack {
	return MinStack{
		dataStack: []int{},
		minStack:  []int{},
	}
}

func (this *MinStack) Push(val int) {
	//首先加入到数据栈的栈顶
	this.dataStack = append(this.dataStack, val)
	//如果小于等于最小栈的栈顶就加入到最小栈
	if len(this.minStack) == 0 || (len(this.minStack) > 0 && val <= this.GetMin()) {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	//从数据栈弹出一个元素
	val := this.dataStack[len(this.dataStack)-1]
	this.dataStack = this.dataStack[:len(this.dataStack)-1]
	//如果等于最小栈栈顶，就从最小栈栈顶弹出
	if val == this.GetMin() {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.dataStack[len(this.dataStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
