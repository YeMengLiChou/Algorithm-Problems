package main

// https://atcoder.jp/contests/abc442/tasks/abc442_c

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type Int = int

func C(n, m int) int {
	d := 1
	for i, j := n, 1; j <= m; i-- {
		d = d * i / j
		j++
	}
	return d
}

func solution() {
	N, M := readI(), readI()
	conflicts := map[int]map[int]bool{}
	for i := range N {
		conflicts[i+1] = map[int]bool{}
	}
	for range M {
		A, B := readI(), readI()
		conflicts[A][B] = true
		conflicts[B][A] = true
	}
	for i := 1; i <= N; i++ {
		c := N - 1 - len(conflicts[i])
		d := 0
		if c >= 3 {
			d = C(c, 3)
		}
		puts(d, " ")
	}
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
