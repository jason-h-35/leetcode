func validMountainArray(arr []int) bool {
    if len(arr) < 3 {
        return false
    }
    state := 0 // 0: look for increase, 1 look for peak, 2look for descend
    for i:=1;i<len(arr);i++ {
        if state == 0 {
            if arr[i-1] >= arr[i] {
                return false
            }
            state = 1
        }
        if state == 1 {
            if arr[i-1] >= arr[i] {
                state = 2
            }
        }
        if state == 2 {
            if arr[i-1] <= arr[i] {
                return false
            }
        }
    }
    return state == 2
}