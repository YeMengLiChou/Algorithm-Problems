package main

import (
	"fmt"
	"math"
)

func main() {
	// input := []int{0, -2, -1, -3, 0, 2, -1}
	// input := []int{1, 4, 2, 7}
	input := []int{198, 883, -791, -540, -453, 238, 897, 411, 599}
	res := maxSumTrionic(input)
	fmt.Println(res)
}

func maxSumTrionic(nums []int) int64 {
	var (
		n         = len(nums)
		i         = 1
		res int64 = math.MinInt64
	)
	for i < n {
		// 第一段递增起点
		if nums[i-1] < nums[i] {
			var (
				l             = i - 1
				lSum    int64 = 0
				lSumMax int64 = math.MinInt64
			)
			for i < n && nums[i-1] < nums[i] {
				lSum += int64(nums[i-1])
				i++
			}
			if i == n || nums[i-1] == nums[i] {
				continue
			}
			// 第二段递减起点
			p := i - 1
			mSum := int64(nums[p])
			for i < n && nums[i-1] > nums[i] {
				mSum += int64(nums[i])
				i++
			}
			if i-1 == p || i == n || nums[i-1] == nums[i] {
				continue
			}
			var (
				q       = i - 1
				rSum    = int64(0)
				rSumMax = int64(math.MinInt64)
			)
			// 第三段递减起点
			for i < n && nums[i-1] < nums[i] {
				rSum += int64(nums[i])
				rSumMax = max(rSumMax, rSum)
				i++
			}
			if i-1 == q {
				continue
			}
			fmt.Println(l, p, q, i-1, lSum, mSum, rSum)
			lSumMax = lSum
			for l < p-1 {
				lSum -= int64(nums[l])
				lSumMax = max(lSumMax, lSum)
				l++
			}
			fmt.Println(lSumMax, rSumMax)
			res = max(res, lSumMax+mSum+rSumMax)
			i = q
			continue
		}
		i++
	}
	return res
}
