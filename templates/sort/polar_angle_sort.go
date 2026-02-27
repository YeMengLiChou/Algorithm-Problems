// Package sort
package sort

import "sort"

type Vec struct{ x, y int }

func Cross(a, b Vec) int {
	return a.x*b.y - a.y*b.x
}

func Half(v Vec) int {
	if v.y > 0 || (v.y == 0 && v.x > 0) {
		return 0 // 上半平面
	}
	return 1 // 下半平面
}

// PolarAngleSort 极角排序
func PolarAngleSort(vs []Vec) {
	sort.Slice(vs, func(i, j int) bool {
		a, b := vs[i], vs[j]
		ha, hb := Half(a), Half(b)
		// 先判断是否处于同一个半平面
		if ha != hb {
			return ha < hb
		}
		// 再按叉积
		c := Cross(a, b)
		if c != 0 { // c > 0 时，a 角度更小
			return c > 0
		}
		// 相同角下，一般按照距离大小排序
		la, lb := a.x*a.x+a.y*a.y, b.x*b.x+b.y*b.y
		if la != lb {
			return la < lb
		}
		if a.x != b.x {
			return a.x < b.x
		}
		return a.y < b.y
	})
}
