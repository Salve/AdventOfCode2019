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

func Test_shortestDist(t *testing.T) {
	type args struct {
		a        string
		b        string
		orbiting map[string]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example1",
			args{
				"YOU",
				"SAN",
				map[string]string{
					"B":   "COM",
					"C":   "B",
					"D":   "C",
					"E":   "D",
					"F":   "E",
					"G":   "B",
					"H":   "G",
					"I":   "D",
					"J":   "E",
					"K":   "J",
					"L":   "K",
					"YOU": "K",
					"SAN": "I",
				},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDist(tt.args.a, tt.args.b, tt.args.orbiting); got != tt.want {
				t.Errorf("shortestDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
