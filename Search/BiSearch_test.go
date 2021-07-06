package search

import (
	"testing"
)

var (
	a  = []int{2, 3, 4, 4, 6, 7, 8, 12, 90}
	a1 = []int{}
)

func TestBiSearch(t *testing.T) {
	type args struct {
		array []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "pt-1", args: args{array: a, k: 4}, want: 2},
		{name: "pt-2", args: args{array: a1, k: 4}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiSearch(tt.args.array, tt.args.k); got != tt.want {
				t.Errorf("BiSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiSearchTheFirst(t *testing.T) {
	type args struct {
		array []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "pt-1", args: args{array: a, k: 4}, want: 2},
		{name: "pt-2", args: args{array: a1, k: 4}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiSearchTheFirst(tt.args.array, tt.args.k); got != tt.want {
				t.Errorf("BiSearchTheFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiSearchTheLast(t *testing.T) {
	type args struct {
		array []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "pt-1", args: args{array: a, k: 4}, want: 3},
		{name: "pt-2", args: args{array: a1, k: 4}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiSearchTheLast(tt.args.array, tt.args.k); got != tt.want {
				t.Errorf("BiSearchTheLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiSearchTheFirstLarger(t *testing.T) {
	type args struct {
		array []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "pt-1", args: args{array: a, k: 4}, want: 4},
		{name: "pt-2", args: args{array: a1, k: 4}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BiSearchTheFirstLarger(tt.args.array, tt.args.k); got != tt.want {
				t.Errorf("BiSearchTheFirstLarger() = %v, want %v", got, tt.want)
			}
		})
	}
}
