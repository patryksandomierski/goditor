package morestrings

import (
	"fmt"
	"testing"
)

func TestIntMin(t *testing.T) {
	ans := ReverseChars("abc")
	expected := "cba"
	if ans != expected {
		t.Errorf("expected: %s\n got: %s", expected, ans)
	}
}

func TestIntMinDataDriver(t *testing.T) {
	var tests = []struct {
		a        string
		expected string
	}{
		{"asd", "dsa"},
		{"123456", "654321"},
		{"aaaaaa", "aaaaaa"},
	}

	for _, td := range tests {
		testname := fmt.Sprintf("%s,%s", td.a, td.expected)
		t.Run(testname, func(t *testing.T) {
			ans := ReverseChars(td.a)
			if ans != td.expected {
				t.Errorf("expected: %s\n got: %s", td.expected, ans)
			}
		})
	}
}
