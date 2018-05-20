package main

// HasDupeChars will return false if string contains duplicate char
func HasDupeChars(s string) bool {
	counter := map[string]int{}

	for _, r := range s {
		counter[string(r)]++
	}

	for k := range counter {
		if counter[k] > 1 {
			return true
		}
	}

	return false
}
