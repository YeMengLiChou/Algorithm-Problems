package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)

	volumes := 0
	playing := false

	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a == 1 {
			volumes += 1
		} else if a == 2 {
			if volumes > 0 {
				volumes -= 1
			}
		} else if a == 3 {
			playing = !playing
		}
		if playing && volumes >= 3 {
			out.WriteString("Yes\n")
		} else {
			out.WriteString("No\n")
		}
	}
}
