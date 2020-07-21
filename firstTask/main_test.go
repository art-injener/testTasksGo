package main

import (
	"reflect"
	"testing"
)

func TestArrayProcessor(t *testing.T) {

	var tests = []struct {
		name string
		A    []int
		B    []int
		AA   []int
		AB   []int
		BB   []int
	}{
		{"compare",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{6, 7, 8, 9, 10, 11, 12, 13, 14},
			[]int{1, 2, 3, 4, 5},
			[]int{6, 7, 8, 9, 10},
			[]int{11, 12, 13, 14},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA, gotAB, gotB := ArrayProcessor(tt.A, tt.B); !reflect.DeepEqual(gotA, tt.AA) ||
				!reflect.DeepEqual(gotAB, tt.AB) ||
				!reflect.DeepEqual(gotB, tt.BB) {
				t.Errorf("\nArrayProcessor() :\n AA = %v ,\n AB = %v,\n BB = %v,\n"+
					"want :\n AA = %v ,\n AB = %v,\n BB = %v\n", gotA, gotAB, gotB, tt.AA, tt.AB, tt.BB)
			}

		})
	}
}
