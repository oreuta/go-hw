package main

import (
	"reflect"
	"testing"
)

var flattenTests = []struct {
	in  [][]int
	out []int
}{
	{
		[][]int{{1, 2}, {3, 4}},
		[]int{1, 2, 4, 3},
	},
	{
		[][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
		[]int{1, 2, 3, 3, 3, 2, 1, 1, 2},
	},
}

func TestFlatten(t *testing.T) {
	for _, v := range flattenTests {
		res := FlattenSlice(v.in)
		if !reflect.DeepEqual(res, v.out) {
			t.Errorf("Flatten(%v)\n Have: %v; \n Want: %v", v.in, res, v.out)
		}
	}
}
