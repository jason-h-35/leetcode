import "sort"
// O(n log n)
func heightChecker(heights []int) int {
    count := 0 // O(1)
    expected := make([]int, len(heights)) // O(n)
    copy(expected, heights) // O(n)
    sort.Ints(expected) // O(n lg n)
    for i:=0;i<len(heights);i++ { // O(n)
        if heights[i] != expected[i] {
            count++
        }
    }
    return count
}
