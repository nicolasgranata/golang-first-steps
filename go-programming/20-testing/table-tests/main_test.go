package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data []int
		answer int
	}

	tests := []test{
		test{[]int{5, 5}, 10},
		test{[]int{10, 10}, 20},
		test{[]int{15, 15}, 30},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}