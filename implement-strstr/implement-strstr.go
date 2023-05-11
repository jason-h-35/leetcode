// scrub through h. when h[i] == n[i] check if the rest matches.
// if so, return the initial index of it.
func strStr(haystack string, needle string) int {
    h := []rune(haystack)
    n := []rune(needle)
    for i:=0; i<len(h); i++ {
        // if needle is bigger than remaining haystack, it won't work.
        if len(h) - i < len(n) {
            return -1
        }
        // found first matching element
        if h[i] == n[0] {
            // save result while we play with i
            result := i
            // scrub through needle 
            for j:=0; j<len(n); j++ {
                // j is the offset added to both i for h and 0 for j
                if h[i+j] != n[0+j] {
                    break
                }
                if j == len(n) - 1 {
                    return result
                }
            }
        }
    }
    return -1
}