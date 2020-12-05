package main

import (
	"strings"
	"testing"
)

type HillTest struct {
	hill string
	want int
}

var hillTests = []HillTest{
	{".....\n.....\n.....", 0},
	{"...#.\n.#.#.\n..#..", 1},
	{".....\n...#.\n.#...", 2},
	{"#####\n#####\n#####", 2},
	{".....\n...#.\n.#...\n.#...\n..#..\n##...", 4},
}

/**
0
.....
.....
.....
**/

/**
1
...#.
.#.#.
..#..
**/

/**
2
.....
...#.
.#...
**/

/**
3
.....
...#.
.#...
.#...
..#..
##...
**/

func TestCountTrees(t *testing.T) {
	for i, test := range hillTests {
		got := CountTrees(strings.Split(test.hill, "\n"))
		if got != test.want {
			t.Errorf("Input(%v) is %v; want %v", i, got, test.want)
		}
	}
}
