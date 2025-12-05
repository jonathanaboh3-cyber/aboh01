package piscine

func Split(s, sep string) []string {
	var result []string
	sepLen := len(sep)
	word := ""

	for i := 0; i < len(s); {
		// Check if the separator starts here
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			// If there’s a word collected, add it to result
			result = append(result, word)
			word = ""
			i += sepLen // skip the separator
		} else {
			word += string(s[i])
			i++
		}
	}

	// Add the last part (if not empty or even if empty after last separator)
	result = append(result, word)

	return result
}
