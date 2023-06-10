func firstPalindrome(words []string) string {
wloop:
	for _, w := range words {
		word := []rune(w)
		for i, j := 0, len(word)-1; i <= j; i, j = i+1, j-1 {
			if word[i] != word[j] {
				continue wloop
			}
		}
		return w
	}
	return ""
}
