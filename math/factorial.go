/*
What is a factorial? I think an example is better than an explaination in
this case:

Factorial(4)
4 * Factorial(3)
4 * 3 * Factorial(2)
4 * 3 * 2 * Factorial(1)
4 * 3 * 2 * 1
*/

package math

// FactorialIterative count factorial by iterative method
func FactorialIterative(num int) int {
	factorial, current := 1, 1
	for current < num {
		current++
		factorial *= current
	}
	return factorial
}

// FactorialRecursive count factorial by recursive method
func FactorialRecursive(num int) int {
	if num == 1 {
		return num
	}
	return num * FactorialIterative(num-1)
}
