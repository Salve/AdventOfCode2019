package main

import (
	"reflect"
	"testing"
)

func Test_runInOut(t *testing.T) {
	type args struct {
		p     []int
		input int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int
	}{
		{
			"example1",
			args{[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, 0},
			[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			"example2",
			args{[]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, 0},
			[]int{1219070632396864},
		},
		{
			"example3",
			args{[]int{104, 1125899906842624, 99}, 0},
			[]int{1125899906842624},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := runInOut(tt.args.p, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("runInOut() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
