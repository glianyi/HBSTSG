package blind

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minStock := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-minStock > maxProfit {
			maxProfit = prices[i] - minStock
		}
		if prices[i] < minStock {
			minStock = prices[i]
		}
	}
	return maxProfit
}
