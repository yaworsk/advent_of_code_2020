package main

import "testing"

type Input struct {
	input string
	want  int
}

var rowTests = []Input{
	{"BFFFBBFRRR", 70},
	{"FFFBBBFRRR", 14},
	{"BBFFBBFRLL", 102},
}

var colTests = []Input{
	{"BFFFBBFRRR", 7},
	{"FFFBBBFRRR", 7},
	{"BBFFBBFRLL", 4},
}

func TestGetRow(t *testing.T) {
	for _, test := range rowTests {
		got := GetRow(test.input[0:7])
		if got != test.want {
			t.Errorf("Input(%s) is %v; want %v)", test.input, got, test.want)
		}
	}
}

func TestGetCol(t *testing.T) {
	for _, test := range colTests {
		got := GetCol(test.input[7:10])
		if got != test.want {
			t.Errorf("Input(%s) is %v; want %v)", test.input, got, test.want)
		}
	}
}
