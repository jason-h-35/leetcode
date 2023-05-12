
func maxProfit(prices []int) int {
    maxProfit := 0
    minPrice := math.MaxInt32
    for _, price := range prices {
        // smallest value seen so far
        if price < minPrice {
            minPrice = price
        // for the smallest value we've seen so far, does this value give higher profit.
        } else if potential := price - minPrice; potential > maxProfit {
            maxProfit = potential
        }
    }
    return maxProfit
}
