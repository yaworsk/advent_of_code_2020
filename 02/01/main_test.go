package main

import "testing"

type PasswordCheckTest struct {
	input PasswordInput
	want  bool
}

var passwordTests = []PasswordCheckTest{
	PasswordCheckTest{
		input: PasswordInput{
			min:      1,
			max:      3,
			letter:   "a",
			password: "abcde",
		},
		want: true,
	},
	PasswordCheckTest{
		input: PasswordInput{
			min:      2,
			max:      5,
			letter:   "b",
			password: "cdefg",
		},
		want: false,
	},
	PasswordCheckTest{
		input: PasswordInput{
			min:      2,
			max:      9,
			letter:   "c",
			password: "ccccccccc",
		},
		want: true,
	},
}

func TestPasswordCheck(t *testing.T) {
	for _, test := range passwordTests {
		got := PasswordCheck(test.input)
		if got != test.want {
			t.Errorf("Input(%s) is %v; want %v", test.input.password, got, test.want)
		}
	}
}
