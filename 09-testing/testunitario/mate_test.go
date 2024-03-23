package testunitario

import "testing"

func TestSum(t *testing.T) {
	table := []struct {
		a, b, expected int
	}{
		{5, 5, 10},
		{2, 3, 5},
		{0, 0, 0},
	}

	for _, test := range table {
		total := Sum(test.a, test.b)
		if total != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, total)
		}
	}
}

func TestGetMax(t *testing.T) {
	table := []struct {
		a, b, expected int
	}{
		{5, 6, 6},
		{4, 3, 4},
		{0, 0, 0},
	}

	for _, test := range table {
		max := GetMax(test.a, test.b)
		if max != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, max)
		}
	}
}

func TestFibonacci(t *testing.T) {
	table := []struct {
		n, expected int
	}{
		{1, 1},
		{10, 55},
		{15, 610},
		{40, 102334155},
	}

	for _, test := range table {
		fib := Fibonacci(test.n)
		if fib != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, fib)
		}
	}
}