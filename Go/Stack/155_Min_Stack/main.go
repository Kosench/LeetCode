package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)
	if len(m.minStack) == 0 || val <= m.minStack[len(m.minStack)-1] {
		m.minStack = append(m.minStack, val)
	}
}

func (m *MinStack) Pop() {
	if len(m.stack) == 0 {
		return
	}

	top := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]

	if m.minStack[len(m.minStack)-1] == top {
		m.minStack = m.minStack[:len(m.minStack)-1]
	}
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.minStack[len(m.minStack)-1]
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
