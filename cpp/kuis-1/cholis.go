package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("challange start ...")

	const inc = 3
	var ptr, res int
	var skip bool = true

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if skip {
			skip = false
			continue
		}

		str := scanner.Text()
		idx := ptr + inc
		if idx > (len(str) - 1) {
			idx = (idx - (len(str) - 1)) - 1
		}

		sym := string(str[idx])
		if sym == "#" {
			res += 1
		}
		ptr = idx
	}

	fmt.Println("res:", res)
}
