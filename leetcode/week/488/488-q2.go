package main

func mergeAdjacent(nums []int) []int64 {
	n := len(nums)
	stack := make([]int64, n)
	cnt := 0
	for i := range n {
		x := int64(nums[i])
		if cnt > 0 && x == stack[cnt-1] {
			stack[cnt-1] = x << 1
			for cnt > 1 && stack[cnt-2] == stack[cnt-1] {
				stack[cnt-2] = stack[cnt-1] << 1
				cnt--
			}
		} else {
			stack[cnt] = x
			cnt++
		}
	}
	res := append(make([]int64, 0), stack[:cnt]...)
	return res
}
