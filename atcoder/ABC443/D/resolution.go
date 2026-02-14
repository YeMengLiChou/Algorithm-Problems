package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func solution() {
	N := int(readI())
	R := make([]int32, N)
	for i := range N {
		R[i] = int32(readI())
	}
	lefts, rights := make([]int32, N), make([]int32, N)
	lefts[0] = R[0]

	for i := range N - 1 {
		lefts[i+1] = min(lefts[i]+1, R[i+1])
	}
	rights[N-1] = lefts[N-1]
	for i := N - 2; i >= 0; i-- {
		rights[i] = min(lefts[i], rights[i+1]+1)
	}
	var res int64 = 0
	for i := range N {
		res += int64(R[i] - rights[i])
	}
	puts(res, "\n")
}

var (
	file = os.Getenv("LI_INPUT")
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
	T := readI()
	for T > 0 {
		solution()
		T--
	}
}

func readI() int64 { // 快读int64
	var x int64
	var neg bool = false
	var c byte
	var err error
	for c, err = in.ReadByte(); c < '0' || c > '9'; c, err = in.ReadByte() {
		if c == '-' {
			neg = true
		}
		if err == io.EOF {
			return 0
		}
	}
	for c >= '0' && c <= '9' {
		x = x*10 + int64(c-'0')
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

