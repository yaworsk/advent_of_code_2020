package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PasswordInput struct {
	min      int
	max      int
	letter   string
	password string
}

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	result := 0

	//Read the input
	for scanner.Scan() {
		line := scanner.Text()
		input := FormatPassword(line)
		if PasswordCheck(input) {
			result += 1
		}
	}

	fmt.Println(result)
}

func FormatPassword(i string) PasswordInput {
	parts := strings.Split(i, " ")
	nums := strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(nums[0])
	max, _ := strconv.Atoi(nums[1])

	return PasswordInput{
		min,
		max,
		strings.TrimSuffix(parts[1], ":"),
		parts[2],
	}
}

func PasswordCheck(p PasswordInput) bool {
	fLetter := p.password[p.min-1]
	sLetter := p.password[p.max-1]
	letter := p.letter[0]

	if fLetter == letter && sLetter != letter {
		return true
	} else if fLetter != letter && sLetter == letter {
		return true
	} else {
		return false
	}
}
