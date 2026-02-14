package main

import (
	"fmt"
	"strconv"
)

func main() {
	var line string
	fmt.Scanln(&line)
	N, _ := strconv.Atoi(line)

	volumes := 0
	playing := false
	for range N {
		fmt.Scanln(&line)
		a, _ := strconv.Atoi(line)
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
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

import (
	"fmt"
	"strconv"
)

func main() {
	var line string
	fmt.Scanln(&line)
	N, _ := strconv.Atoi(line)

	volumes := 0
	playing := false
	for range N {
		fmt.Scanln(&line)
		a, _ := strconv.Atoi(line)
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
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

import "fmt"

func main() {
	var line string
	fmt.Scanln(&line)

	ans := 0
	for _, x := range line {
		if x == 'i' || x == 'j' {
			ans++
		}
	}
	fmt.Println(ans)
}
