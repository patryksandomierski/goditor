package morestrings

import (
	"testing"
)

func TestIntMin(t *testing.T) {
	ans := ReverseChars("abc")
	expected := "cba"
	if ans != expected {
		t.Errorf("expected: %s\n got: %s", expected, ans)
	}
}