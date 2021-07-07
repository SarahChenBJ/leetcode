package search

import "testing"

func TestSmallestDivisor(t *testing.T) {
	type args struct {
		array     []int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "pt-1", args: args{array: []int{2, 3, 5, 7, 11}, threshold: 11}, want: 3},
		{name: "pt-2", args: args{array: []int{19}, threshold: 5}, want: 4},
		{name: "pt-3", args: args{array: []int{1, 2, 5, 9}, threshold: 6}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestDivisor(tt.args.array, tt.args.threshold); got != tt.want {
				t.Errorf("SmallestDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPeak(t *testing.T) {
	type args struct {
		array     []int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "pt-1", args: args{array: []int{1, 2, 1, 3, 5, 6, 7, 4, 5}}, want: 6},
		{name: "pt-2", args: args{array: []int{1, 2, 3, 4, 5, 6, 7}}, want: 6},
		{name: "pt-3", args: args{array: []int{825}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPeak(tt.args.array); got != tt.want {
				t.Errorf("SmallestDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchInRotatedArray(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "pt-1", args: args{array: []int{4, 5, 6, 7, 8, 1, 2, 3, 4}, target: 2}, want: 6},
		{name: "pt-2", args: args{array: []int{}, target: 5}, want: -1},
		{name: "pt-3", args: args{array: []int{1, 2, 3, 9}, target: 9}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInRotatedArray(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("SmallestDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}
