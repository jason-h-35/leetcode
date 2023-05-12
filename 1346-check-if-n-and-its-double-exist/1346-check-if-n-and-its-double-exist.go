func checkIfExist(arr []int) bool {
    for i, numi := range arr {
        for j, numj := range arr {
            if numi == 2 * numj && i != j {
                return true
            }
        }
    }
    return false
}