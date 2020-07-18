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
		{name: "test 1", args: args{n: 89, p: 1}, want: 1},
		{name: "test 2", args: struct {
			n int
			p int
		}{n: 92, p: 1}, want: -1},
		{name: "test 2", args: args{n: 46288, p: 3}, want: 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DigPow(tt.args.n, tt.args.p); got != tt.want {
				t.Errorf("DigPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompArr(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				A: []int{121, 144, 19, 161, 19, 144, 19, 11},
				B: []int{121, 14641, 20736, 361, 25921, 361, 20736, 361},
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				A: []int{121, 144, 19, 161, 19, 144, 19, 11},
				B: []int{132, 14641, 20736, 361, 25921, 361, 20736, 361},
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				A: nil,
				B: []int{132, 14641, 20736, 361, 25921, 361, 20736, 361},
			},
			want: false,
		},
		{
			name: "test 4",
			args: args{
				A: []int{121, 144, 19, 161, 19, 144, 19},
				B: []int{132, 14641, 20736, 361, 25921, 361, 20736, 361},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if eq := CompArr(tt.args.A, tt.args.B); eq != tt.want {
				t.Errorf("CompArr() = %v, want %v", eq, tt.want)
			}
		})
	}
}
