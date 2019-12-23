package main

import (
	"reflect"
	"testing"
)

func Test_fft(t *testing.T) {
	type args struct {
		in    []int
		phase int
		p     [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"example1", args{in: []int{1, 2, 3, 4, 5, 6, 7, 8}, phase: 4, p: patterns(8)}, []int{0, 1, 0, 2, 9, 4, 9, 8}},
		{"example2", args{in: []int{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5}, phase: 100, p: patterns(32)}, []int{2, 4, 1, 7, 6, 1, 7, 6}},
		{"example3", args{in: []int{1, 9, 6, 1, 7, 8, 0, 4, 2, 0, 7, 2, 0, 2, 2, 0, 9, 1, 4, 4, 9, 1, 6, 0, 4, 4, 1, 8, 9, 9, 1, 7}, phase: 100, p: patterns(32)}, []int{7, 3, 7, 4, 5, 4, 1, 8}},
		{"example4", args{in: []int{6, 9, 3, 1, 7, 1, 6, 3, 4, 9, 2, 9, 4, 8, 6, 0, 6, 3, 3, 5, 9, 9, 5, 9, 2, 4, 3, 1, 9, 8, 7, 3}, phase: 100, p: patterns(32)}, []int{5, 2, 4, 3, 2, 1, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fft(tt.args.in, tt.args.phase, tt.args.p)[:8]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fft() = %v, want %v", got, tt.want)
			}
		})
	}
}
