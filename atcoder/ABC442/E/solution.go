package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

type Int = int

func solution() {
	N, Q := readI(), readI()
	// 记录原始编号与 id  的映射
	raws := make([]int, N)
	// 统计同一个方向上的点数
	cnts := make([]int, 0, N)
	dirs := make([]Dir, 0, N)
	// dir -> id
	dirIDs := make(map[Dir]int, N)
	for i := range N {
		x, y := readI(), readI()
		o := GCD(abs(x), abs(y))
		d := Dir{x / o, y / o}
		id, ok := dirIDs[d]
		if !ok {
			id = len(dirs)
			dirIDs[d] = id
			cnts = append(cnts, 0)
			dirs = append(dirs, d)
		}
		cnts[id]++
		raws[i] = id
	}
	dirCnt := len(dirs)
	orders := make([]int, dirCnt)
	for i := range orders {
		orders[i] = i
	}
	// 极角排序
	sort.Slice(orders, func(i int, j int) bool {
		a, b := dirs[orders[i]], dirs[orders[j]]
		ah, bh := Half(a), Half(b)
		if ah != bh { // 先按半平面排序
			return ah < bh
		}
		// 再按叉积排序
		v := Cross(a, b)
		if v != 0 {
			return v > 0
		}
		// 最后按x，y比较，但因为上面归一化后，基本走不到
		if a.x != b.x {
			return a.x < b.x
		}
		return a.y < b.y
	})

	// 方便根据 id 找到排序后的位置
	pos := make([]int, dirCnt)
	for i, id := range orders {
		pos[id] = i
	}
	// 根据排序结果，处理出前缀和
	s := make([]int, dirCnt+1)
	for i := range dirCnt {
		s[i+1] = s[i] + cnts[orders[i]]
	}
	for ; Q > 0; Q-- {
		// raw id
		a, b := readI()-1, readI()-1
		// map id
		aid, bid := raws[a], raws[b]
		// 转换为排序后的位置
		ap, bp := pos[aid], pos[bid]

		// 相同角
		if ap == bp {
			puts(cnts[aid], "\n")
		} else if bp < ap {
			puts(s[ap+1]-s[bp], "\n")
		} else {
			puts(s[dirCnt]-s[bp]+s[ap+1], "\n")
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GCD(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Dir struct{ x, y int }

func Half(d Dir) int {
	if d.y > 0 || (d.y == 0 && d.x > 0) {
		return 0 // 上半平面
	}
	return 1 // 下半平面
}

func Cross(a, b Dir) int {
	return a.x*b.y - b.x*a.y
}

var (
	file = os.Getenv("input")
	in   = bufio.NewReader(os.Stdin)
	out  = bufio.NewWriter(os.Stdout)
)

func main() {
	if len(file) > 0 {
		_, err := os.Stat(file)
		out.WriteString("(use file)\n")
		if err == nil {
			f, _ := os.Open(file)
			in = bufio.NewReader(f)
		}
	}
	defer out.Flush()
	T := 1
	// T = readI()
	for range T {
		solution()
	}
}

func readI() Int { // 快读
	var (
		x   Int
		neg bool
		c   byte
		err error
	)
	for c, err = in.ReadByte(); c < '0' || c > '9'; c, err = in.ReadByte() {
		if c == '-' {
			neg = true
		}
		if err == io.EOF {
			return 0
		}
	}
	for c >= '0' && c <= '9' {
		x = x*10 + Int(c-'0')
		c, _ = in.ReadByte()
	}
	if neg {
		return -x
	}
	return x
}

func puts(x ...any) { // 快写
	for _, val := range x {
		fmt.Fprint(out, val)
	}
}

func readS() string { // 快读字符串
	var strBuffer bytes.Buffer
	var c byte
	var err error
	for {
		c, err = in.ReadByte()
		if err != nil {
			if err == io.EOF {
				return ""
			}
			return ""
		}
		if c == ' ' || c == '\n' || c == '\r' || c == '\t' {
			continue
		}
		break
	}
	strBuffer.WriteByte(c)
	for {
		c, err = in.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if c == ' ' || c == '\n' || c == '\r' || c == '\t' {
			in.UnreadByte()
			break
		}
		strBuffer.WriteByte(c)
	}
	return strBuffer.String()
}
