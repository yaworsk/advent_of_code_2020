package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Passport struct {
	byr, iyr, eyr, pid, hgt, hcl, ecl, cid string
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	vs := 0
	p := Passport{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if CheckValid(p) {
				vs++
			}
			p = Passport{}
			continue
		}

		attrs := strings.Split(line, " ")
		for _, v := range attrs {
			attr := strings.Split(v, ":")
			switch attr[0] {
			case "byr":
				p.byr = attr[1]
			case "iyr":
				p.iyr = attr[1]
			case "eyr":
				p.eyr = attr[1]
			case "pid":
				p.pid = attr[1]
			case "hgt":
				p.hgt = attr[1]
			case "hcl":
				p.hcl = attr[1]
			case "ecl":
				p.ecl = attr[1]
			case "cid":
				p.cid = attr[1]
			}
		}
	}
	if CheckValid(p) {
		vs++
	}
	fmt.Println(vs)
}

func CheckValid(p Passport) bool {
	if p.byr == "" || p.iyr == "" || p.eyr == "" || p.pid == "" || p.hgt == "" || p.hcl == "" || p.ecl == "" {
		return false
	}

	return true
}
