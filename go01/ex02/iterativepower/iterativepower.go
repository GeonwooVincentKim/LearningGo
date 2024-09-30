package iterativepower

func IterativePower(nb int, power int) int {
	result := 1
	for power > 0 {
		result *= nb
		power--
	}
	return result
}
