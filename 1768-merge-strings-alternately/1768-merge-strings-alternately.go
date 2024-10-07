import "strings"
import "unicode/utf8"

func mergeAlternately(word1 string, word2 string) string {
  var b strings.Builder
  i, j := 0, 0
  for i < len(word1) && j < len(word2) {
    // alternate with Builder
    r, size := utf8.DecodeRuneInString(word1[i:])
    b.WriteRune(r)
    i += size
    r, size = utf8.DecodeRuneInString(word2[j:])
    b.WriteRune(r)
    j += size
  }
  // figure out the one with extra and WriteString
  if i < len(word1) {
    b.WriteString(word1[i:])
  }
  if j < len(word2) {
    b.WriteString(word2[j:])
  }
  // Bulider -> String
  return b.String()
}