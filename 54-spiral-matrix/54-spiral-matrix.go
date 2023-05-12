

func spiralOrder(matrix [][]int) []int {
    DIRECTIONS := map[int][]int{ 
        0: []int{0,1},
        1: []int{1,0},
        2: []int{0,-1},
        3: []int{-1,0},
    }
    result := make([]int, len(matrix) * len(matrix[0]))
    // create boolean matrix of traversed values
    seen := make([][]bool, len(matrix))
    for i, _ := range matrix {
        seen[i] = make([]bool, len(matrix[i]))
    }
    // select a starting point and direction
    i, j := 0, 0
    leni, lenj := len(matrix), len(matrix[0])
    dir, count := 0, 0
    // grab, move
    for count < len(result) {
        // grab
        result[count] = matrix[i][j]
        seen[i][j] = true
        count += 1
        // move
        ti, tj := i, j
        ti += DIRECTIONS[dir][0]; tj += DIRECTIONS[dir][1]
        // if seen or index not in range, change direction, and apply i, j
        if !InInterval(ti, leni) || !InInterval (tj, lenj) {
            dir = (dir + 1) % len(DIRECTIONS)
            i += DIRECTIONS[dir][0]
            j += DIRECTIONS[dir][1]
        } else if seen[ti][tj] {
            dir = (dir + 1) % len(DIRECTIONS)
            i += DIRECTIONS[dir][0]
            j += DIRECTIONS[dir][1]
        } else {
            i, j = ti, tj
        }
    }
    return result
}

func InInterval(index int, length int) bool {
    return 0 <= index && index < length
}