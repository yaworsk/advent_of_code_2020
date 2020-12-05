package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr, iyr, eyr           int
	pid, hgt, hcl, ecl, cid string
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
				num, err := strconv.Atoi(attr[1])
				if err != nil {
					p.byr = 0
				} else {
					p.byr = num
				}
			case "iyr":
				num, err := strconv.Atoi(attr[1])
				if err != nil {
					p.iyr = 0
				} else {
					p.iyr = num
				}
			case "eyr":
				num, err := strconv.Atoi(attr[1])
				if err != nil {
					p.eyr = 0
				} else {
					p.eyr = num
				}
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
	/**
	byr (Birth Year) - four digits; at least 1920 and at most 2002.
	iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	hgt (Height) - a number followed by either cm or in:
	If cm, the number must be at least 150 and at most 193.
	If in, the number must be at least 59 and at most 76.
	hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	pid (Passport ID) - a nine-digit number, including leading zeroes.
	cid (Country ID) - ignored, missing or not.
	**/

	if p.byr < 1920 || p.byr > 2002 {
		return false
	}

	if p.iyr < 2010 || p.iyr > 2020 {
		return false
	}

	if p.eyr < 2020 || p.eyr > 2030 {
		return false
	}

	end := len(p.hgt) - 2
	if end <= 0 {
		return false
	}
	if p.hgt[end:] == "cm" {
		h, err := strconv.Atoi(p.hgt[0:end])
		if err != nil {
			return false
		}
		if h < 150 || h > 193 {
			return false
		}
	} else if p.hgt[end:] == "in" {
		h, err := strconv.Atoi(p.hgt[0:end])
		if err != nil {
			return false
		}
		if h < 59 || h > 76 {
			return false
		}
	} else {
		return false
	}

	if len(p.hcl) != 7 {
		return false
	}
	if string(p.hcl[0]) != "#" {
		return false
	}
	matched, _ := regexp.Match(`^[0-9a-f]+$`, []byte(p.hcl[1:6]))
	if !matched {
		return false
	}

	if p.ecl != "amb" && p.ecl != "blu" && p.ecl != "brn" && p.ecl != "gry" && p.ecl != "grn" && p.ecl != "hzl" && p.ecl != "oth" {
		return false
	}

	if len(p.pid) != 9 {
		return false
	}

	matched, _ = regexp.Match(`^[0-9]+$`, []byte(p.pid))
	if !matched {
		return false
	}

	return true
}
