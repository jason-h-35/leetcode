func getRow(numRows int) []int {
    result := make([][]int, numRows+1)
    // O(n)
    for i:=0;i<len(result);i++ {
        result[i] = make([]int, i+1)
        for j:=0;j<len(result[i]);j++ {
            if  j==0 || i==j {
                result[i][j] = 1
            } else {
                result[i][j] = result[i-1][j] + result[i-1][j-1]
            }
        }
    }
    return result[len(result)-1]
}