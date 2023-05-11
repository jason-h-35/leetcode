func replaceElements(arr []int) []int {
    m := arr[len(arr)-1]
    arr[len(arr)-1] = -1
    for i:=len(arr)-2; 0 <= i; i-- {
        tmp := arr[i]
        arr[i] = m
        if tmp > m {
            m = tmp
        }
    }
    return arr
}