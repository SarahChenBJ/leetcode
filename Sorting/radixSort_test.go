package sorting

import (
	"reflect"
	"testing"
)

func TestRadixSortLSD(t *testing.T) {

	tests := []struct {
		name  string
		array []string
		want  []string
	}{
		{name: "pt-1", array: []string{"Auger", "Abh", "Fanrong", "Changsheng", "Sarah", "hahah"}, want: []string{"Abh", "Auger", "Changsheng", "Fanrong", "hahah", "Sarah"}},
		{name: "pt-2", array: []string{"Auger", "", "Sarah", ""}, want: []string{"", "", "Auger", "Sarah"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if RadixSortLSD(tt.array); !reflect.DeepEqual(tt.array, tt.want) {
				t.Errorf("output:%v, want:%v", tt.array, tt.want)
			}
		})
	}
}
