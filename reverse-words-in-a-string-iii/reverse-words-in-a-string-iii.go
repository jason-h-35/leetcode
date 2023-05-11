import "strings"
func reverseWords(s string) string {
    words := strings.Split(s, " ")
    for i, word := range words {
        words[i] = reverse(word)
    }
    return strings.Join(words, " ")
}

func reverse(s string) string {
    var b strings.Builder
    for i:=len(s)-1;i>=0;i-- {
        b.WriteByte(s[i])
    }
    return b.String()
}