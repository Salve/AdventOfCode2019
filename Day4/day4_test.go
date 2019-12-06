package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{111111}, true},
		{"example2", args{223450}, false},
		{"example3", args{123789}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.v); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValid2(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{112233}, true},
		{"example2", args{123444}, false},
		{"example3", args{111122}, true},
		{"test1", args{911222}, false},
		{"test2", args{133444}, true},
		{"test3", args{999999}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid2(tt.args.v); got != tt.want {
				t.Errorf("isValid2() = %v, want %v", got, tt.want)
			}
		})
	}
}
