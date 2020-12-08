package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	qs := ""
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			if !strings.Contains(qs, string(line[i])) {
				qs += string(line[i])
			}
		}

		if line == "" {
			total += len(qs)
			qs = ""
		}
	}
	total += len(qs)
	fmt.Println("Sum: ", total)
}
