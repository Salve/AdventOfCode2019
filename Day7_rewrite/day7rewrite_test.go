package main

import "testing"

func Test_tryFeedbackSequence(t *testing.T) {
	type args struct {
		seq []int
		prg []int
	}
	tests := []struct {
		name       string
		args       args
		wantSignal int
	}{
		{"part2example1", args{
			seq: []int{9, 8, 7, 6, 5},
			prg: []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
		}, 139629729},
		{"part2example2", args{
			seq: []int{9, 7, 8, 5, 6},
			prg: []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
		}, 18216},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSignal := tryFeedbackSequence(tt.args.seq, tt.args.prg); gotSignal != tt.wantSignal {
				t.Errorf("tryFeedbackSequence() = %v, want %v", gotSignal, tt.wantSignal)
			}
		})
	}
}
