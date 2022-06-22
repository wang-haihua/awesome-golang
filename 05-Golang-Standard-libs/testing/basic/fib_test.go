package basic

import "testing"

// 单 case 测试
/*
func TestFib(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)
	actual := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}
*/

// 多 case 测试
func TestFib(t *testing.T) {
	var testCases = []struct {
		in       int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, tc := range testCases {
		actual := Fib(tc.in)
		if actual != tc.expected {
			t.Errorf("Fib(%d) = %d; expected %d", tc.in, actual, tc.expected)
		}
	}
}
