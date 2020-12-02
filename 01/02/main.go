package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	CheckError(err)

	scanner := bufio.NewScanner(f)
	pos := 0
	var entries = make([]int, 200)

	//Read the input
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		CheckError(err)
		entries[pos] = n
		pos++
	}

	//Check answer
	for _, v := range entries {
		for _, v2 := range entries {
			for _, v3 := range entries {
				if v+v2+v3 == 2020 {
					fmt.Println("Answer: ", v*v2*v3)
					os.Exit(1)
				}
			}
		}
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
