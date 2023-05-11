import "strconv"
func findNumbers(nums []int) int {
    count := 0
    for _, num := range nums {
        if isEven(num) {
            count++
        }
    }
    return count
}

func isEven(num int) bool {
    return len(strconv.Itoa(num)) % 2 == 0
}