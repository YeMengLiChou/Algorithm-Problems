package main

func longestBalancedV1(s *string) int {
	n := len(*s)
	mp := make(map[byte]int)
	res := 1
	checkSame := func() bool {
		cnt := 0
		for _, v := range mp {
			if cnt == 0 {
				cnt = v
			} else if cnt != v {
				return false
			}
		}
		return true
	}

	for i := range n {
		clear(mp)
		for j := i; j < n; j++ {
			mp[(*s)[j]]++
			if checkSame() {
				res = max(res, j-i+1)
			}
		}
	}
	return res
}

func longestBalanced(s string) int {
	return longestBalancedV1(&s)
}
