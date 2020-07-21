package main

import "testing"

func Test_checkBrackets(t *testing.T) {

	tests := []struct {
		testString string
		want       bool
	}{
		{"((b)", false},
		{")(", false},
		{"{(a+b)*[c-d]}", true},
		{"((b)", false},
		{"(a[0]+b[2c[6]]) {24 + 53}", true},
		{"([1-2)]", false},
		{"([1-2]))", false},
		{"([1-2])", true},
	}
	for _, tt := range tests {
		t.Run(tt.testString, func(t *testing.T) {
			if got := checkBrackets(tt.testString); got != tt.want {
				t.Errorf("checkBrackets() = %v, want %v", got, tt.want)
			}
		})
	}
}
