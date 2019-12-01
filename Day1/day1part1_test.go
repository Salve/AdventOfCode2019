package main

import "testing"

func Test_fuelRequired(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{mass: 1969}, 654},
		{"example2", args{mass: 100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fuelRequired(tt.args.mass); got != tt.want {
				t.Errorf("fuelRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
