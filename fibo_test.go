package main

import "testing"

func TestFibonacci(t *testing.T) {
	type fibonacciTestcase struct {
		input int
		want  int
	}

	var testcases = []fibonacciTestcase{
		{
			4,
			5,
		},
		{
			5,
			8,
		},
		{
			6,
			13,
		},
		{
			25,
			121393,
		},
	}

	for _, c := range testcases {
		got := fibonacci(c.input)
		if got != c.want {
			t.Errorf("Incorrect return value from fibonacci for input '%v'; wanted '%v', got '%v'", c.input, c.want, got)
		}
	}
}
