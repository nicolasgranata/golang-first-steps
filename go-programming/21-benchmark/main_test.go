package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}

func BenchmarkMySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mySum(2, 3)
	}
}