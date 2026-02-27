package templates

// 排列

func A(n, m int) int {
	d := int(1)
	for i := n - m + 1; i <= m; i++ {
		d *= i
	}
	return d
}

// 组合

func C(n, m int) int {
	d := int(1)
	for i, j := n, 1; j <= n-m; i-- {
		d = d * i / j
		j++
	}
	return d
}

// 最大公约数

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func GCD1(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

// 最小公倍数

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
