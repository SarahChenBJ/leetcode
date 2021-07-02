package swap

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{name: "normal", array: []int{2, 0, 2, 1, 1, 0}, want: []int{0, 0, 1, 1, 2, 2}},
		{name: "normal-2", array: []int{2, 1, 2, 0, 0, 1}, want: []int{0, 0, 1, 1, 2, 2}},
		{name: "zero", array: []int{0}, want: []int{0}},
		{name: "one", array: []int{1}, want: []int{1}},
		{name: "two", array: []int{2}, want: []int{2}},
		{name: "empty", array: []int{}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if SortColors(tt.array); !reflect.DeepEqual(tt.array, tt.want) {
				t.Errorf("SortColors = %v, want: %v ", tt.array, tt.want)
			}
		})
	}
}
