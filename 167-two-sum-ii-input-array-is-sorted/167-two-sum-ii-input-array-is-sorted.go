func twoSum(numbers []int, target int) []int {
    for i, num1 := range numbers {
        for j, num2 := range numbers {
            if num1 + num2 == target && i != j {
                if j < i {
                    i, j = j, i
                }
                return []int{i+1, j+1}
            }
        }
    }
    return nil
}