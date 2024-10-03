package iterativefactorial

// Init Statement: Executed before the first iteration
// Condition Expression: Evaluated before every iteration
// The Post statement: Executed at the end of every iteration
func IterativeFactorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
