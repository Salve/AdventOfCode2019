package main

import "testing"

func Test_fuelRequired(t *testing.T) {
	type args struct {
		mass      int
		recursive bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"nonrecursive_1", args{mass: 1969, recursive: false}, 654},
		{"nonrecursive_2", args{mass: 100756, recursive: false}, 33583},
		{"recursive_1", args{mass: 14, recursive: true}, 2},
		{"recursive_2", args{mass: 1969, recursive: true}, 966},
		{"recursive_3", args{mass: 100756, recursive: true}, 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fuelRequired(tt.args.mass, tt.args.recursive); got != tt.want {
				t.Errorf("fuelRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
