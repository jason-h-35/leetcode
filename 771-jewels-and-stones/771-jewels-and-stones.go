func numJewelsInStones(jewels string, stones string) int {
    jewelMap := make(map[rune]bool)
    for _, j := range jewels {
        jewelMap[j] = true
    }
    numJewels := 0
    for _, s := range stones {
        if jewelMap[s] {
            numJewels++
        }
    }
    return numJewels
}