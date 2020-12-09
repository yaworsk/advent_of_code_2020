package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BagHier struct {
	bag     string
	parents []string
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	var hier = make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " contain")

		parent := StripBags(s[0])
		children := s[1:]

		for _, v := range strings.Split(children[0], ",") {
			k := CleanString(v)
			hier[k] = append(hier[k], parent)
		}
	}

	fmt.Println(CountParents("shiny gold", hier))
}

func StripBags(s string) string {
	return s[:strings.Index(s, " bag")]
}

func CleanString(s string) string {
	c := s[3:]
	if string(c[len(c)-1]) == "." {
		c = c[:len(c)-1]
	}
	return StripBags(c)
}

func CountParents(s string, h map[string][]string) int {
	fmt.Println(s, h[s])
	for _, v := range h[s] {
		if _, ok := h[v]; ok {
			CountParents(v, h)
		} else {
			fmt.Println(v)
			continue
		}
	}
	return 1
}
