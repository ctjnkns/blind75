func maxProfit(prices []int) int {
    l, r := 0, 0
    max := 0
    for r < len(prices) {
        total := prices[r] - prices[l]
        if total > 0 {
            if total > max {
                max = total
            }
        } else {
            l = r
        }
        r++
    }
    return max
}
