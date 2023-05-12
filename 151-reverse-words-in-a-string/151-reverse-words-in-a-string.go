import "strings"
func reverseWords(s string) string {
    words := strings.Split(s, " ")
    for l, r := 0, len(words) - 1; l < r; l, r = l+1, r-1 {
        words[l], words[r] = words[r], words[l]
    }
    n:=0
    for _, word := range words {
        if word != "" {
            words[n] = word
            n++
        }
    }
    words = words[:n]
    return strings.Join(words, " ")
}