package main

import (
	"fmt"
)

func minimumPairRemoval(nums []int) int {
	size := len(nums)
	ans := 0
	MAX_VALUE := int(^uint(0) >> 1)
	for size > 1 {
		minValue := MAX_VALUE
		minIdx := 0
		isSorted := true
		for idx := range size - 1 {
			if nums[idx] > nums[idx+1] {
				isSorted = false
			}
			sum := nums[idx] + nums[idx+1]
			if sum < minValue {
				minIdx = idx
				minValue = sum
			}
		}
		if isSorted {
			break
		}
		nums[minIdx] = minValue
		nums = append(nums[:minIdx+1], nums[minIdx+2:]...)
		size--
		ans++
	}
	return ans
}

func main() {
	nums := make([]int, 0)
	nums = append(nums, 1000, 999, 998, 997, 996, 995, 994, 993, 992, 991, 990, 989, 988, 987, 986, 985, 984, 983, 982, 981, 980, 979, 978, 977, 976, 975, 974, 973, 972, 971, 970, 969, 968, 967, 966, 965, 964, 963, 962, 961, 960, 959, 958, 957, 956, 955, 954, 953, 952, 951)

	// nums = append(nums, 2, 2, -1, 3, -2, 2, 1, 1, 1, 0, -1)

	// nums = append(nums, -2, 1, 2, -1, -1, -2, -2, -1, -1, 1, 1)

	result := minimumPairRemoval(nums)
	fmt.Println(result)
}
