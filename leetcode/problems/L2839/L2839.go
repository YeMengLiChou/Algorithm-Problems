package main

func canBeEqual(s1 string, s2 string) bool {
	odd, even := make([]int, 26), make([]int, 26)
	for i, x := range s1 {
		x -= 'a'
		if i&1 == 1 {
			odd[x]++
		} else {
			even[x]++
		}
	}

	for i, x := range s2 {
		x -= 'a'
		if i&1 == 1 {
			odd[x]--
		} else {
			even[x]--
		}
	}
	for i := range 26 {
		if odd[i] != 0 || even[i] != 0 {
			return false
		}
	}
	return true
}
