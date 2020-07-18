package codewar

import (
	"reflect"
	"testing"
)

func TestShortestBinary(t *testing.T) {
	t.Skip()
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{name: "test 1", args: args{
			A: []int{0, 1, 1, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1},
			B: []int{0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1}},
			want: []int{0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortestBinary(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortestBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
