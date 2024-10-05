package basicatoi2

func BasicAtoi2(s string) int {
	result := 0

	for _, char := range s {
		// Check if the character is not a digit (outside '0' to '9')
		if char < '0' || char > '9' {
			return 0
		}

		// Convert character to digit by subtracting '0' and update the result
		digit := int(char - '0')
		result = result*10 + digit
	}

	return result
}
