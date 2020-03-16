package codewar

import "testing"

func TestSmallestInt(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test 1", args: struct{ A []int }{A: []int{1, 3, 6, 4, 1, 2}}, want: 5},
		{name: "test 2", args: struct{ A []int }{A: []int{1, 2, 3}}, want: 4},
		{name: "test 3", args: struct{ A []int }{A: []int{-1, -3}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestInt(tt.args.A); got != tt.want {
				t.Errorf("SmallestInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
