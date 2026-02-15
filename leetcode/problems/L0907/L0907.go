package main

import "fmt"

func main() {
	arr := []int{3, 1, 2, 4}
	res := sumSubarrayMins(arr)
	fmt.Println(res)
}

func sumSubarrayMins(arr []int) int {
	const M int = 1e9 + 7
	arr = append(arr, -1)
	n := len(arr)
	stack := make([]int, n)
	cnt := 0
	res := 0
	for i := range n {
		x := arr[i]
		// 单调栈，单调递增
		for cnt > 0 && arr[stack[cnt-1]] > x {
			// 栈顶元素弹出时，说明 i 是其连续子数组的右边界
			top := stack[cnt-1]
			cnt--
			// 栈顶第二个元素是其左边界
			left := -1
			if cnt > 0 {
				left = stack[cnt-1]
			}
			fmt.Println(left, top, i, arr[top])
			res = (res + (top-left)*(i-top)*arr[top]) % M

		}
		stack[cnt] = i
		cnt++
	}
	return res
}
