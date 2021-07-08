package BST

import (
	"reflect"
	"testing"

	BiTree "github.com/sarahchen/leetcode/BiTree/Traversal"
)

func TestInsertBST(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "pt-1", args: args{array: []int{9, 6, 12, 3, 10, 16, 13, 19, 101}}, want: []int{3, 6, 9, 10, 12, 13, 16, 19, 101}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertBST(tt.args.array)
			preorder := BiTree.Preorder(got)
			if !reflect.DeepEqual(preorder, tt.want) {
				t.Errorf("InsertBST() = %v, want %v", preorder, tt.want)
			}
		})
	}
}

func TestBalancedBSTFromAscArray(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "pt-1", args: args{array: []int{1, 2, 3, 5, 7, 12}}, want: []int{1, 2, 3, 5, 7, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BalancedBSTFromAscArray(tt.args.array, 0, len(tt.args.array)-1)
			preorder := BiTree.Preorder(got)
			if !reflect.DeepEqual(preorder, tt.want) {
				t.Errorf("BalancedBSTFromAscArray() = %v, want %v", preorder, tt.want)
			}
		})
	}
}
