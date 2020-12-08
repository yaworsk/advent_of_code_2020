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
	qs2 := ""
	start := true
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		if start {
			for i := 0; i < len(line); i++ {
				if !strings.Contains(qs, string(line[i])) {
					qs += string(line[i])
					qs2 += string(line[i])
				}
			}
			start = false
		} else if line == "" {
			total += len(qs)
			qs = ""
			qs2 = ""
			start = true
		} else {
			for i, p := 0, 0; i < len(qs); i++ {
				if strings.Contains(line, string(qs[i])) {
					p++
				} else {
					qs2 = qs2[:p] + qs[i+1:]
				}
			}
			qs = qs2
		}
	}
	total += len(qs)
	fmt.Println("Sum: ", total)
}
