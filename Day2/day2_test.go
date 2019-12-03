package main

import (
	"reflect"
	"testing"
)

func Test_runIntcode(t *testing.T) {
	type args struct {
		program []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"example1", args{program: []int{1, 0, 0, 0, 99}}, []int{2, 0, 0, 0, 99}},
		{"example2", args{program: []int{2, 3, 0, 3, 99}}, []int{2, 3, 0, 6, 99}},
		{"example3", args{program: []int{2, 4, 4, 5, 99, 0}}, []int{2, 4, 4, 5, 99, 9801}},
		{"example4", args{program: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runIntcode(tt.args.program); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runIntcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
