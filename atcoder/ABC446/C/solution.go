package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func solution() {
	N, D := readI(), readI()
	dq := make([]int, N+1)
	hh, tt := 0, 0
	for range N {
		dq[tt] = readI()
		tt++
	}
	for i := range N {
		x := readI()
		for hh <= i && x > 0 {
			if dq[hh] >= x {
				dq[hh] -= x
				x = 0
			} else {
				x -= dq[hh]
				hh++
			}
		}
		for i-hh >= D {
			hh++
		}
	}
	res := 0
	for hh <= tt {
		res += dq[hh]
		hh++
	}
	puts(res, "\n")
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
	T := Int(1)
	T = readI()
	for range T {
		solution()
	}
}

type Int = int

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
