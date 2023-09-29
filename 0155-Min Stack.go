package main

type MinStack struct {
    top *StackNode
    min int 
}

type StackNode struct {
    data int
    next *StackNode
    lastMin int
}


func Constructor() MinStack {
    return MinStack {
        top : nil,
    }
}


func (this *MinStack) Push(val int)  {
    var newTop *StackNode
    if this.top == nil {
        newTop = &StackNode{
            data : val,
            next : nil,
            lastMin : val,
        }
    } else {
        newTop = &StackNode{
            data : val,
            next : this.top,
        }
        if this.top.lastMin < val{
            newTop.lastMin = this.top.lastMin
        } else {
            newTop.lastMin = val
        }
    }
    this.top = newTop
    this.min = this.top.lastMin
}


func (this *MinStack) Pop()  {
    if this.top.next == nil {
        this.top = nil
        this.min = -9999999999
    } else {
        this.top = this.top.next
        this.min = this.top.lastMin
    }
}


func (this *MinStack) Top() int {
    return this.top.data
}


func (this *MinStack) GetMin() int {
    return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
