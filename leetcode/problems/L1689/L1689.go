package main

func minPartitions(n string) int {
	mx := '0'
	for _, x := range n {
		mx = max(mx, x)
	}
	return int(mx - '0')
}
