package main

import (
	"strconv"
	"testing"

	"golang.org/x/tour/tree"
)

func TestSame(t *testing.T) {
	tests := []struct {
		same   bool
		values [2]int
	}{
		{values: [2]int{1, 1}, same: true},
		{values: [2]int{2, 2}, same: true},
		{values: [2]int{10, 10}, same: true},
		{values: [2]int{1, 2}, same: false},
		{values: [2]int{2, 3}, same: false},
		{values: [2]int{10, 12}, same: false},
	}

	for _, test := range tests {
		name := "{" + strconv.Itoa(test.values[0]) + "," + strconv.Itoa(test.values[1]) + "} should be " + strconv.FormatBool(test.same)
		t.Run(name, func(t *testing.T) {
			isSame := Same(
				tree.New(test.values[0]),
				tree.New(test.values[1]),
			)

			if isSame != test.same {
				t.Errorf("values: %v should be %v, but is %v", test.values, test.same, isSame)
			}
		})
	}
}
