package ex01

import "sort"

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

func minCoins2(val int, coins []int) []int {
	var res1, res2 []int
	var sum1 int

	for _, c := range coins {
		if c <= 0 || len(coins) == 0 {
			return []int{}
		}
	}
	sort.Ints(coins)
	res := make([]int, 0, len(coins))
	temp := map[int]struct{}{}
	for _, item := range coins {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			res = append(res, item)
		}
	}

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
