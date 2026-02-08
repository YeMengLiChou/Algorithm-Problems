package main

func minimumDeletions(s string) int {
	leftB, rightA := 0, 0
	for _, c := range s {
		if c == 'a' {
			rightA++
		}
	}
	res := leftB + rightA
	for _, c := range s {
		if c == 'a' {
			rightA--
		} else {
			leftB++
		}
		res = min(res, leftB+rightA)
	}
	return res
}
