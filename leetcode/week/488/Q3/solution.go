package main

func countSubarrays(nums []int, k int64) int64 {
	n := len(nums)
	res := int64(0)
	mindq := make([]int, 0, n)
	maxdq := make([]int, 0, n)

	l := 0
	for r := 0; r < n; r++ {
		for len(maxdq) > 0 && nums[maxdq[(len(maxdq)-1)]] <= nums[r] {
			maxdq = maxdq[:len(maxdq)-1]
		}
		maxdq = append(maxdq, r)

		for len(mindq) > 0 && nums[mindq[len(mindq)-1]] >= nums[r] {
			mindq = mindq[:len(mindq)-1]
		}
		mindq = append(mindq, r)

		for l <= r {
			mx := int64(nums[maxdq[0]])
			mn := int64(nums[mindq[0]])
			length := int64(r - l + 1)
			if (mx-mn)*length <= k {
				break
			}
			if maxdq[0] == l {
				maxdq = maxdq[1:]
			}
			if mindq[0] == l {
				mindq = mindq[1:]
			}
			l++
		}
		res += int64(r - l + 1)
	}
	return res
}
