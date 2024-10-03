package recursivefactorial

func RecursiveFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * RecursiveFactorial(n-1)
}
