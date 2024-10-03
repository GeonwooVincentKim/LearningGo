package findnextprime

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}

	for candidate := nb; ; candidate++ {
		isPrime := true
		for divisor := 2; divisor*divisor <= candidate; divisor++ {
			if candidate%divisor == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return candidate
		}
	}
}
