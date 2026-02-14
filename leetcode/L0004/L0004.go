package main

func findMedianSortedArraysV1(nums1 []int, nums2 []int) float64 {
	const INF int = 1e7
	n, m := len(nums1), len(nums2)
	nums := make([]int, n+m)
	idx := 0
	for i, j := 0, 0; i < n || j < m; {
		x, y := INF, INF
		if i < n {
			x = nums[i]
		}
		if j < m {
			y = nums[j]
		}
		if x <= y {
			nums[idx] = x
			idx++
			i++
		} else {
			nums[idx] = y
			idx++
			j++
		}
	}
	mid := (n + m) >> 1
	if (n+m)&1 == 0 {
		return (float64(nums[mid-1]) + float64(nums[mid])) / 2
	}
	return float64(nums[mid])
}

func findMedianSortedArraysV2(nums1 []int, nums2 []int) float64 {
	const INF int = 1e7
	n, m := len(nums1), len(nums2)
	i, j := 0, 0
	mid := (n + m) >> 1
	isEven := (n+m)&1 == 0
	pre := 0
	for i+j <= mid {
		x, y := INF, INF
		if i < n {
			x = nums1[i]
		}
		if j < m {
			y = nums2[j]
		}
		if i+j == mid {
			if isEven {
				return (float64(pre) + float64(min(x, y))) / 2
			} else {
				return float64(min(x, y))
			}
		}
		if x <= y {
			i++
			pre = x
		} else {
			j++
			pre = y
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArraysV1(nums1, nums2)
}
