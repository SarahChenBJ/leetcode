package BiTree

import (
	"reflect"
	"testing"

	"github.com/sarahchen/leetcode/dataStruct"
)

var (
	btree = &dataStruct.TreeNode{
		Val: 0,
		Left: &dataStruct.TreeNode{
			Val:  1,
			Left: &dataStruct.TreeNode{Val: 3},
		},
		Right: &dataStruct.TreeNode{
			Val:  2,
			Left: &dataStruct.TreeNode{Val: 4},
			Right: &dataStruct.TreeNode{
				Val:   5,
				Right: &dataStruct.TreeNode{Val: 6},
			},
		},
	}
	btree_level    = []int{0, 1, 2, 3, 4, 5, 6}
	btree_preorder = []int{3, 1, 0, 4, 2, 5, 6}
)

func TestPreorder(t *testing.T) {
	type args struct {
		node *dataStruct.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "pt-1", args: args{node: btree}, want: btree_preorder},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Preorder(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
