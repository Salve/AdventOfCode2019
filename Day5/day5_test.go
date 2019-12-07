package main

import (
	"reflect"
	"testing"
)

func Test_runIntcode(t *testing.T) {
	type args struct {
		program []int
		input   int
	}

	diagnosticprogram := readIntcode("input")

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"part1", args{diagnosticprogram, 1}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 9431221}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runIntcode(tt.args.program, tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runIntcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
