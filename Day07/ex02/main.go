package ex02

// minCoins2 accepts a necessary amount and a sorted slice of unique denominations of coins.
//It may be something like [1,5,10,50,100,500,1000] or something exotic, like [1,3,4,7,13,15].
//The output is supposed to be a slice of coins of minimal size that can be used to express the value
//(e.g. for 13 and [1,5,10] it should give you [10,1,1,1])

import "sort"

func minCoins2(val int, coins []int) []int {
	var res1, res2 []int
	var sum1 int

	// checks negative values of coins slice, so it will not have endless loop
	for _, c := range coins {
		if c <= 0 || len(coins) == 0 {
			return []int{}
		}
	}

	//sort coins slice at first, so the algorithm starts getting coins from the biggest denomination
	sort.Ints(coins)

	//go through coins slice starting from different coin to find better solution
	res := make([]int, 0, len(coins))
	temp := map[int]struct{}{}
	for _, item := range coins {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			res = append(res, item)
		}
	}
	//checks sum of result coins to fit given value
	//In cycle call MinCoins without lastelem in sclice, check results and save if this correct and have less elements
	for i := 0; i < len(coins); i++ {
		res2 = minCoins(val, coins[:len(coins)-i])
		if countSum(res2) == val {
			if len(res1) == 0 {
				res1 = res2
			} else if len(res1) > len(res2) {
				res1 = res2
			}
		}
	}
	for _, i := range res1 {
		sum1 += i
	}
	if sum1 == val {
		return res1
	}
	return []int{}
}

func countSum(coins []int) (sum int) {
	for _, i := range coins {
		sum += i
	}
	return
}

func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}
