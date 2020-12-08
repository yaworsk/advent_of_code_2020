package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	hseat := 0

	//Read the input
	for scanner.Scan() {
		line := scanner.Text()
		r := line[0:7]
		s := line[7:10]
		row := GetRow(r)
		col := GetCol(s)
		sid := row*8 + col
		if sid > hseat {
			hseat = sid
		}
	}
	fmt.Println("Highest seat: ", hseat)
}

func GetRow(s string) int {
	m := 127
	l := 0
	c := ""
	for i := 0; i < len(s)-1; i++ {
		c = string(s[i])
		if c == "F" {
			m = m - ((m - l) / 2) - 1
		} else {
			l = l + ((m - l) / 2) + 1
		}
	}
	if string(s[len(s)-1]) == "F" {
		return l
	} else {
		return m
	}
}

func GetCol(s string) int {
	m := 7
	l := 0
	c := ""
	for i := 0; i < len(s)-1; i++ {
		c = string(s[i])
		if c == "L" {
			m = m - ((m - l) / 2) - 1
		} else {
			l = l + ((m - l) / 2) + 1
		}
	}
	if string(s[len(s)-1]) == "L" {
		return l
	} else {
		return m
	}
}
