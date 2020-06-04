package main

import (
	"reflect"
	"testing"
)

type testGoldenRatio struct {
	values  [][]int
	average []int
}

var testsGoldenRatio = []testGoldenRatio{
	{[][]int{
		[]int{1, 2, 3}, 
		[]int{4, 5, 6}, 
		[]int{7, 8, 9}}, 
		[]int{5, 4, 7, 8, 9, 6, 3, 2}},
}

func TestSplitter(t *testing.T) {
	for _, pair := range testsGoldenRatio {
		v := GoldenRatio(pair.values)
		if !reflect.DeepEqual(v, pair.average) {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
