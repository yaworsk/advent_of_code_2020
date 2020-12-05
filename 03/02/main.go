package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	i, result := 0, 1
	hill := make([]string, 323)

	//Read the input
	for scanner.Scan() {
		line := scanner.Text()
		hill[i] = line
		i++
	}

	steps := []int{1, 3, 5, 7}
	for _, v := range steps {
		result *= CountTrees(hill, v, 1)
	}

	result *= CountTrees(hill, 1, 2)

	fmt.Println(result)
}

func CountTrees(hill []string, h int, v int) int {
	p := 0
	t := 0

	for i, val := range hill {
		if i == 0 {
			continue
		}

		if v == 2 && i%2 == 1 {
			continue
		}

		p += h
		p = p % len(val)

		if string(val[p]) == "#" {
			t++
		}
	}

	return t
}
