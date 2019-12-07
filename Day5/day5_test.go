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

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"part1", args{readIntcode("input"), 1}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 9431221}},

		{"equal8_positional_8", args{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 8}, []int{1}},
		{"equal8_positional_9", args{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 9}, []int{0}},
		{"equal8_positional_-1", args{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, -1}, []int{0}},
		{"equal8_positional_0", args{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 0}, []int{0}},
		{"equal8_positional_9999", args{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 9999}, []int{0}},

		{"equal8_immediate_8", args{[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 8}, []int{1}},
		{"equal8_immediate_9", args{[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 9}, []int{0}},
		{"equal8_immediate_-1", args{[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, -1}, []int{0}},
		{"equal8_immediate_0", args{[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 0}, []int{0}},
		{"equal8_immediate_9999", args{[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 9999}, []int{0}},

		{"less8_positional_7", args{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 7}, []int{1}},
		{"less8_positional_-1", args{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, -1}, []int{1}},
		{"less8_positional_0", args{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 0}, []int{1}},
		{"less8_positional_8", args{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 8}, []int{0}},
		{"less8_positional_9999", args{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 9999}, []int{0}},

		{"less8_immediate_7", args{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 7}, []int{1}},
		{"less8_immediate_-1", args{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, -1}, []int{1}},
		{"less8_immediate_0", args{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 0}, []int{1}},
		{"less8_immediate_8", args{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 8}, []int{0}},
		{"less8_immediate_9999", args{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 9999}, []int{0}},

		{"jmp_positional_0", args{[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 0}, []int{0}},
		{"jmp_positional_8", args{[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 8}, []int{1}},
		{"jmp_positional_9999", args{[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, 9999}, []int{1}},
		{"jmp_positional_-1", args{[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, -1}, []int{1}},

		{"jmp_immediate_0", args{[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 0}, []int{0}},
		{"jmp_immediate_8", args{[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 8}, []int{1}},
		{"jmp_immediate_9999", args{[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, 9999}, []int{1}},
		{"jmp_immediate_-1", args{[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, -1}, []int{1}},

		{"full_0", args{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 0}, []int{999}},
		{"full_-1", args{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, -1}, []int{999}},
		{"full_7", args{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 7}, []int{999}},
		{"full_8", args{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 8}, []int{1000}},
		{"full_9", args{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 9}, []int{1001}},
		{"full_9999", args{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 9999}, []int{1001}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runIntcode(tt.args.program, tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runIntcode() = %v, want %v", got, tt.want)
			}
		})
	}
}