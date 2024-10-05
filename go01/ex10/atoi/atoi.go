package atoi

// Atoi converts a string to an integer, handling optional '+' or '-' signs.
// Returns 0 if the string contains non-digit characters (ignoring leading sign).
func Atoi(s string) int {
	// Initialize variables
	result := 0
	sign := 1
	start := 0

	// Handle optional '+' or '-' sign at the beginning
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		if s[0] == '-' {
			sign = -1
		}
		start = 1 // Skip the sign for further processing
	}

	// Iterate over the string starting from the character after the sign
	for i := start; i < len(s); i++ {
		// Check if the character is not a digit
		if s[i] < '0' || s[i] > '9' {
			return 0
		}

		// Convert character to digit by subtracting '0' and update the result
		digit := int(s[i] - '0')
		result = result*10 + digit
	}

	// Apply the sign to the result
	return result * sign
}
