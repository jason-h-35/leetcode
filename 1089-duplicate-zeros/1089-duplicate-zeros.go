// O(n^2)
func duplicateZeros(arr []int)  {
    for i:=0;i<len(arr);i++ {
        if arr[i] == 0 {
            insert(&arr, 0, i) // O(n)
            i++
        }
    }
}

// O(n)
func insert(arr *[]int, num int, index int) {
    aref := *arr
    for i:=len(aref)-1; index < i; i-- {
        aref[i] = aref[i-1]
    }
    aref[index] = num
}