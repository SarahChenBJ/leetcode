package sorting

import (
	"fmt"
	"reflect"
	"testing"

	swap "github.com/sarahchen/leetcode/Sorting/Swap"
	"github.com/sarahchen/leetcode/dataStruct"
	"github.com/sarahchen/leetcode/sorting/insertion"
)

var (
	a1    = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	a1Asc = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a2    = []int{1, 2, 5, 9, 2, 9, 7}
	a2Asc = []int{1, 2, 2, 5, 7, 9, 9}
	a3    = []int{1}
	a3Asc = []int{1}
	a4    = []int{}
	a4Asc = []int{}
)

func convertArrayToLink(a []int) *dataStruct.ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &dataStruct.ListNode{Val: a[0], Next: nil}
	p := head
	for _, v := range a[1:] {
		n := &dataStruct.ListNode{Val: v, Next: nil}
		p.Next = n
		p = n
	}
	return head
}

func IsTheSameLink(l1, l2 *dataStruct.ListNode) bool {
	if l1 == l2 {
		return true
	}
	n1, n2 := l1, l2
	for n1 != nil && n2 != nil && n1.Val == n2.Val {
		n1 = n1.Next
		n2 = n2.Next
		fmt.Println(n1, n2)

	}
	return n1 == nil && n2 == nil
}

func TestDirectInsertionSort(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{name: "direct-1", array: a1, want: a1Asc},
		{name: "repeat", array: a2, want: a2Asc},
		{name: "one", array: a3, want: a3Asc},
		{name: "empty", array: a4, want: a4Asc},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.array
			if insertion.DirectInsertionSort(a); !reflect.DeepEqual(a, tt.want) {
				t.Errorf("DirectInsertionSort() = %v, want %v", a, tt.want)
			}
		})
	}
}

func TestBiInsertionSort(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{name: "direct-1", array: a1, want: a1Asc},
		{name: "repeat", array: a2, want: a2Asc},
		{name: "one", array: a3, want: a3Asc},
		{name: "empty", array: a4, want: a4Asc},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.array
			if insertion.BiInsertionSort(a); !reflect.DeepEqual(a, tt.want) {
				t.Errorf("BiInsertionSort() = %v, want %v", a, tt.want)
			}
		})
	}
}

func TestShellSort(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{name: "direct-1", array: a1, want: a1Asc},
		{name: "repeat", array: a2, want: a2Asc},
		{name: "one", array: a3, want: a3Asc},
		{name: "empty", array: a4, want: a4Asc},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.array
			if insertion.ShellSort(a); !reflect.DeepEqual(a, tt.want) {
				t.Errorf("ShellSort() = %v, want %v", a, tt.want)
			}
		})
	}
}

func TestLinkInsertionSort(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{name: "direct-1", array: a1, want: a1Asc},
		{name: "repeat", array: a2, want: a2Asc},
		{name: "one", array: a3, want: a3Asc},
		{name: "empty", array: a4, want: a4Asc},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := convertArrayToLink(tt.array)
			wantLink := convertArrayToLink(tt.want)
			if l := insertion.LinkInsertionSort(link); !IsTheSameLink(l, wantLink) {
				t.Errorf("ShellSort() = %v, want %v", l, wantLink)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{name: "direct-1", array: a1, want: a1Asc},
		{name: "repeat", array: a2, want: a2Asc},
		{name: "one", array: a3, want: a3Asc},
		{name: "empty", array: a4, want: a4Asc},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.array
			if swap.BubbleSort(a); !reflect.DeepEqual(a, tt.want) {
				t.Errorf("ShellSort() = %v, want %v", a, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		want  []int
	}{
		{name: "direct-1", array: a1, want: a1Asc},
		{name: "repeat", array: a2, want: a2Asc},
		{name: "one", array: a3, want: a3Asc},
		{name: "empty", array: a4, want: a4Asc},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.array
			if swap.QuickSort(a); !reflect.DeepEqual(a, tt.want) {
				t.Errorf("ShellSort() = %v, want %v", a, tt.want)
			}
		})
	}
}
