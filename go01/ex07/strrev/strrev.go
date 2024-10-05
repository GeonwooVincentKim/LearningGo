package strrev

func GetStrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func StrRev(s string) string {
	runes := []rune(s)
	getLength := GetStrLen(s)

	for i := 0; i < getLength/2; i++ {
		runes[i], runes[getLength-i-1] = runes[getLength-i-1], runes[i]
	}

	return string(runes)
}
