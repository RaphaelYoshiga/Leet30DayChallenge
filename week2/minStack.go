package main;

type MinStack struct {
	min int;
	items []int;
}

const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(MaxUint >> 1)

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack { items: []int {}, min: MaxInt};
}

func (this *MinStack) Push(x int)  {
	this.items = append(this.items, x);

	if this.min > x {
		this.min = x;
	}
}

func (this *MinStack) Pop()  {

	top := this.Top();
	this.items = this.items[0:len(this.items) -1];

	if top == this.min {
		min := MaxInt
		for _, n := range this.items{
			if min > n {
				min = n;
			}
		}		
		this.min = min;
	}
}

func (this *MinStack) Top() int {
    return this.items[len(this.items) - 1];
}

func (this *MinStack) GetMin() int {
      return this.min;
}



/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */