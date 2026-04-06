package main

// 关键在模拟过程中能够得知当前方向上可能会碰撞到的障碍物；

type pair struct{ x, y int }

func robotSim(commands []int, obstacles [][]int) int {
	dirs := [4]pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	blocks := make(map[pair]bool, len(obstacles))
	for _, x := range obstacles {
		blocks[pair{x[0], x[1]}] = true
	}
	idx, x, y, ans := 0, 0, 0, 0
	for _, c := range commands {
		switch c {
		// turn right
		case -1:
			idx = (idx + 1) % 4
		// turn left
		case -2:
			idx = (idx + 3) % 4
		// go straight
		default:
			// 需要判断当前方向上是否有障碍物
			d := dirs[idx]
			for range c {
				if blocks[pair{x + d.x, y + d.y}] {
					break
				}
				x, y = x+d.x, y+d.y
			}
			ans = max(ans, x*x+y*y)
		}
	}
	return ans
}
