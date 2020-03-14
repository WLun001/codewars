package codewar

import "testing"

func TestSquareSum(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{name: "test 1", args: []int{1, 2, 2}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquareSum(tt.args); got != tt.want {
				t.Errorf("SquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDigPow(t *testing.T) {
	type args struct {
		n int
		p int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test 1", args: struct {
			n int
			p int
		}{n: 89, p: 1}, want: 1},
		{name: "test 2", args: struct {
			n int
			p int
		}{n: 92, p: 1}, want: -1},
		{name: "test 2", args: struct {
			n int
			p int
		}{n: 46288, p: 3}, want: 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DigPow(tt.args.n, tt.args.p); got != tt.want {
				t.Errorf("DigPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
