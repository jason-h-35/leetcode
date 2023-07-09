func smallestEvenMultiple(n int) int {
    for i:=n; i <= 300; i *= 2 {
        if i % 2 == 0 {
            return i
        }
    }
    return -1
}