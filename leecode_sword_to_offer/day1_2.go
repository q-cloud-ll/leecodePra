package leecode_sword_to_offer

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64}, // 先放入一个最大元素，后续元素只要比较都比他小，进行赋值即可
	}
}

// Push 放入数据
func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[(len(this.minStack) - 1)]
	this.minStack = append(this.minStack, min(top, x))
}

// Pop 删除数据
func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

// Top 获得头数据
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// Min 获取最小数据
func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
