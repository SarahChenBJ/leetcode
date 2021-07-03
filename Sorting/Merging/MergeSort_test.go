package merging

import (
	"reflect"
	"testing"
)

var (
	a1     = []int{1, 2, 3, 0, 0, 0}
	b1     = []int{2, 5, 6}
	wantA1 = []int{1, 2, 2, 3, 5, 6}

	a2     = []int{0, 0, 0, 0}
	b2     = []int{2}
	wantA2 = []int{2, 0, 0, 0}
)

func TestMergeArrays(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "pt-1", args: args{a1, b1}, want: wantA1},
		{name: "pt-2", args: args{a2, b2}, want: wantA2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if MergeArrays(tt.args.a, tt.args.b); !reflect.DeepEqual(tt.args.a, tt.want) {
				t.Errorf("MergeArrays = %v, want = %v", tt.args.a, tt.want)
			}
		})
	}
}
