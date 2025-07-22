package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	storage []int
}

func NewStack(capacity int) *Stack {
	return &Stack{storage: make([]int, 0, capacity)}
}

func (s *Stack) Push(token int) {
	s.storage = append(s.storage, token)
}

func (s *Stack) Pop() int {
	top := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return top
}

func applyOperator(oper string, a, b int) int {
	switch oper {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("неизвестный оператор: " + oper)
	}
}

var operators = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
}

func evalRPN(tokens []string) int {
	stack := NewStack(len(tokens))
	for _, token := range tokens {
		if operators[token] {
			num2 := stack.Pop()
			num1 := stack.Pop()
			stack.Push(applyOperator(token, num1, num2))
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				return -1
			}
			stack.Push(num)
		}

	}
	return stack.Pop()
}

func main() {
	sl2 := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(sl2))

	sl3 := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(sl3))
}

//type Stack[T any] struct {
//	storage []T
//}
//
//func NewStack[T any](capacity int) *Stack[T] {
//	return &Stack[T]{storage: make([]T, 0, capacity)}
//}
//
//func (s *Stack[T]) Push(token T) {
//	s.storage = append(s.storage, token)
//}
//
//func (s *Stack[T]) Pop() (T, error) {
//	if len(s.storage) == 0 {
//		var zero T
//		return zero, fmt.Errorf("stack is empty")
//	}
//	top := s.storage[len(s.storage)-1]
//	s.storage = s.storage[:len(s.storage)-1]
//	return top, nil
//}
