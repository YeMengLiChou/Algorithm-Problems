package main

func isTrionicV1(nums []int) bool {
	var (
		n   = len(nums)
		cnt = 0
		l   = 0
	)
	for i := 1; i < n; i++ {
		pre, cur := nums[i-1], nums[i]
		if pre == cur {
			return false
		}
		if pre < cur { // increase
			if cnt&1 == 0 {
				continue
			}
			if i-1 == l {
				return false
			}
			l = i - 1
			cnt++
		} else {
			if cnt&1 == 1 {
				continue
			}
			if i-1 == l {
				return false
			}
			l = i - 1
			cnt++
		}

		if cnt > 2 {
			return false
		}
	}
	return cnt == 2
}
