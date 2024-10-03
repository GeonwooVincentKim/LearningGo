package iterativefactorial

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 || nb == 1 {
		return 1
	}

	result := 1
	for i := 2; i <= nb; i++ {
		if result > (1<<31-1)/i {
			return 0
		}
		result *= i
	}

	return result
}
