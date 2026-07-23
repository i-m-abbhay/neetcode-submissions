func maxProfit(prices []int) int {
//[10,1,5,6,7,1]
// 10 1 1 1 1 1
//	0 0 4 5 6 0
//			^
	minSoFar := prices[0]
	maxProfit := 0
	for i:=1;i<len(prices);i++{
		minSoFar = min(minSoFar, prices[i])
		maxProfit = max(maxProfit, prices[i]-minSoFar)
	}
	return maxProfit
}
