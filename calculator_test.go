package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(2, 3)
	expected := 5
	if total != expected {
		t.Errorf("Sum produced incorrect output: %d instead of %d", total, expected)
	}
}
