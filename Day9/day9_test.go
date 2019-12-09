package main

import (
	"reflect"
	"testing"
)

func Test_program_runInOut(t *testing.T) {
	type fields struct {
		intcode []int
		loc     int
		rbase   int
		state   state
		input   []int
		output  int
	}
	type args struct {
		input []int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput []int
	}{
		{
			"example1",
			fields{intcode: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}},
			args{[]int{}},
			[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			"example2",
			fields{intcode: []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}},
			args{[]int{}},
			[]int{1219070632396864},
		},
		{
			"example3",
			fields{intcode: []int{104, 1125899906842624, 99}},
			args{[]int{}},
			[]int{1125899906842624},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &program{
				intcode: tt.fields.intcode,
				loc:     tt.fields.loc,
				rbase:   tt.fields.rbase,
				state:   tt.fields.state,
				input:   tt.fields.input,
				output:  tt.fields.output,
			}
			if gotOutput := p.runInOut(tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("runInOut() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
