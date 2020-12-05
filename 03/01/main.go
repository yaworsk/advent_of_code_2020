package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	i := 0
	hill := make([]string, 323)

	//Read the input
	for scanner.Scan() {
		line := scanner.Text()
		hill[i] = line
		i++
	}

	fmt.Println(CountTrees(hill))
}

func CountTrees(hill []string) int {
	p := 0
	t := 0

	for i, v := range hill {
		if i == 0 {
			continue
		}

		p += 3
		p = p % len(v)

		if string(v[p]) == "#" {
			t++
		}
	}

	return t
}
