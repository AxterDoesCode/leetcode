package main

func maxProfit(prices []int) int {
    l, r := 0, 1
    maxProfit := 0

    for r < len(prices) {
        if prices[l] < prices[r] {
            maxProfit = max2(maxProfit, prices[r] - prices[l])
        } else {
            l = r
        }
        r++
    }
    return maxProfit
}

func max2 (a, b int) int {
    if a < b {
        return b
    } else if a > b {
        return a
    }
    return a
}
