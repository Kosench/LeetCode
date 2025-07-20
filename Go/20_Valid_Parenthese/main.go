package main

import "fmt"

type Stack struct {
	items []byte
}

func (s *Stack) Push(char byte) {
	s.items = append(s.items, char)
}

func (s *Stack) Pop() {
	if len(s.items) > 0 {
		s.items = s.items[:len(s.items)-1]
	}
}

func NewStack(capacity int) Stack {
	return Stack{items: make([]byte, 0, capacity)}
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false // Нечетное количество скобок не может быть валидным
	}

	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := NewStack(len(s) / 2)

	for i := 0; i < len(s); i++ {
		char := s[i]

		if openBracket, isClosing := pairs[char]; isClosing {
			if len(stack.items) == 0 || stack.items[len(stack.items)-1] != openBracket {
				return false
			}
			stack.Pop()
		} else {
			stack.Push(char)
		}
	}

	return len(stack.items) == 0
}

func main() {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"([)]", false},
		{"", true},
		{"((", false},
		{"))", false},
		{"({[]})", true},
		{"({[}])", false},
	}

	fmt.Println("=== Тестирование решения Valid Parentheses ===\n")

	for i, tc := range testCases {
		result1 := isValid(tc.input)

		status1 := "✅"
		if result1 != tc.expected {
			status1 = "❌"
		}

		fmt.Printf("Тест %d:\n", i+1)
		fmt.Printf("  Вход: \"%s\"\n", tc.input)
		fmt.Printf("  Ожидается: %t\n", tc.expected)
		fmt.Printf("  Решение 1: %t %s\n", result1, status1)
		fmt.Println()
	}
}
