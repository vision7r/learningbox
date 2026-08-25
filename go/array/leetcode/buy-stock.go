package leetcode

//problem [https://leetcode.com/problems/best-time-to-buy-and-sell-stock/]

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		// Update minimum buying price seen so far
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			// Calculate profit if sold today and update maxProfit if it's higher
			maxProfit = profit
		}
	}

	return maxProfit

}
