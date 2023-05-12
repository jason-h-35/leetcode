func plusOne(digits []int) []int {
    digits[len(digits)-1] += 1
    for i:=len(digits)-1; i!=0; i-- {
        if digits[i] == 10 {
            digits[i] = 0
            digits[i-1] += 1
        }
    }
    // handle leading 10
    if digits[0] == 10 {
        digits[0] = 0
        result := []int{1}
        result = append(result, digits...)
        return result
    }
    return digits
}