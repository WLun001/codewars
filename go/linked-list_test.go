package codewar

import "testing"

func Test_linkedList(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test 1", args: struct{ A []int }{A: []int{1, 4, -1, 3, 2}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linkedList(tt.args.A); got != tt.want {
				t.Errorf("linkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
