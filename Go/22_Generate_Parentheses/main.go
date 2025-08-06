package main

import "fmt"

func backtrack(result *[]string, current string, open int, close int, max int) {
	// Если набрали нужное количество скобок (2*max)
	if len(current) == max*2 {
		*result = append(*result, current) // Добавляем в результат
		return
	}

	// Правило 1: можно добавить '(', если их ещё меньше max
	if open < max {
		backtrack(result, current+"(", open+1, close, max)
	}

	// Правило 2: можно добавить ')', если их меньше чем '('
	if close < open {
		backtrack(result, current+")", open, close+1, max)
	}
}

func generateParenthesis(n int) []string {
	var result []string             // Здесь будем хранить все правильные комбинации
	backtrack(&result, "", 0, 0, n) // Запускаем генерацию
	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
}
