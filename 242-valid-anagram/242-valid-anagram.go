func isAnagram(s string, t string) bool {
    Counts := make(map[rune]int)
    for _, letter := range s {
        Counts[letter] += 1
    }
    for _, letter := range t {
        Counts[letter] -= 1
    }
    // check map equality
    for _, count := range Counts {
        if count != 0 {
            return false
        }
    }
    return true
}