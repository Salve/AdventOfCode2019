package main

import (
	"testing"
)

func Test_findClosestIntersect(t *testing.T) {
	type args struct {
		wirelayouts [][]wireSegment
		startingLoc loc
	}
	tests := []struct {
		name            string
		args            args
		wantMinDistance int
	}{
		{
			"example1",
			args{
				[][]wireSegment{
					{{"R"[0], 8}, {"U"[0], 5}, {"L"[0], 5}, {"D"[0], 3}},
					{{"U"[0], 7}, {"R"[0], 6}, {"D"[0], 4}, {"L"[0], 4}},
				},
				loc{0, 0},
			},
			6,
		},
		{
			"example2",
			args{
				[][]wireSegment{
					{{"R"[0], 75}, {"D"[0], 30}, {"R"[0], 83}, {"U"[0], 83}, {"L"[0], 12}, {"D"[0], 49}, {"R"[0], 71}, {"U"[0], 7}, {"L"[0], 72}},
					{{"U"[0], 62}, {"R"[0], 66}, {"U"[0], 55}, {"R"[0], 34}, {"D"[0], 71}, {"R"[0], 55}, {"D"[0], 58}, {"R"[0], 83}},
				},
				loc{0, 0},
			},
			159,
		},
		{
			"example3",
			args{
				[][]wireSegment{
					{{"R"[0], 98}, {"U"[0], 47}, {"R"[0], 26}, {"D"[0], 63}, {"R"[0], 33}, {"U"[0], 87}, {"L"[0], 62}, {"D"[0], 20}, {"R"[0], 33}, {"U"[0], 53}, {"R"[0], 51}},
					{{"U"[0], 98}, {"R"[0], 91}, {"D"[0], 20}, {"R"[0], 16}, {"D"[0], 67}, {"R"[0], 40}, {"U"[0], 7}, {"R"[0], 15}, {"U"[0], 6}, {"R"[0], 7}},
				},
				loc{0, 0},
			},
			135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMinDistance := findClosestIntersect(tt.args.wirelayouts, tt.args.startingLoc); gotMinDistance != tt.wantMinDistance {
				t.Errorf("findClosestIntersect() = %v, want %v", gotMinDistance, tt.wantMinDistance)
			}
		})
	}
}
