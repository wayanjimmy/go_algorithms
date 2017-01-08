package math

// FactorialIterative Count factorial by iterative method
func FactorialIterative(num int) int {
	factorial, current := 1, 1
	for current < num {
		current++
		factorial *= current
	}
	return factorial
}
