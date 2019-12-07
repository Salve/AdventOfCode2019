package main

import "testing"

func Test_countOrbits(t *testing.T) {
	type args struct {
		orbiting map[string]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{map[string]string{
			"B": "COM",
			"C": "B",
			"D": "C",
			"E": "D",
			"F": "E",
			"G": "B",
			"H": "G",
			"I": "D",
			"J": "E",
			"K": "J",
			"L": "K",
		}}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOrbits(tt.args.orbiting); got != tt.want {
				t.Errorf("countOrbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
