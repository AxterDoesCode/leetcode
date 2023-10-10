package main

type StackNode struct {
    val int
    lastMin int
    nextNode *StackNode
}
type MinStack struct {
    top *StackNode
}


func Constructor() MinStack {
    return MinStack{top : nil}
}


func (this *MinStack) Push(val int)  {
    //Need to calculate the lastMin here
    if this.top == nil {
        this.top = &StackNode{
            val : val, 
            lastMin : val, 
            nextNode : nil,
        }
    } else {
        this.top = &StackNode {
            val : val,
            lastMin : min(val, this.top.lastMin),
            nextNode : this.top,
        }
    }
}


func (this *MinStack) Pop()  {
    this.top = this.top.nextNode
}


func (this *MinStack) Top() int {
    return this.top.val
}


func (this *MinStack) GetMin() int {
    return this.top.lastMin
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
