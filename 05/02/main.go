package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	seats := []int{}

	//Read the input
	for scanner.Scan() {
		line := scanner.Text()
		r := line[0:7]
		s := line[7:10]
		row := GetRow(r)
		col := GetCol(s)
		sid := row*8 + col
		seats = append(seats, sid)
	}
	sort.Ints(seats)

	for i, v := range seats {
		if i == 0 {
			continue
		}
		if v-seats[i-1] == 2 {
			fmt.Println("Your seat: ", v-1)
			os.Exit(1)
		}
	}

	fmt.Println("Seats: ", seats)
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
