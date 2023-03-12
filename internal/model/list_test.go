package model

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_NewListFromArray(t *testing.T) {
	tests := []struct {
		nums []int
		want *ListNode
	}{
		{
			nums: []int{2, 7, 11},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  11,
						Next: nil,
					},
				},
			},
		},
		{
			nums: []int{2},
			want: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
		{
			nums: []int{},
			want: nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(st *testing.T) {

			if got := NewListFromArray(tt.nums); !reflect.DeepEqual(got, tt.want) {
				st.Errorf("given %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ToArray(t *testing.T) {
	tests := []struct {
		nums *ListNode
		want []int
	}{
		{
			nums: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  11,
						Next: nil,
					},
				},
			},
			want: []int{2, 7, 11},
		},
		{
			nums: &ListNode{
				Val:  2,
				Next: nil,
			},
			want: []int{2},
		},
		{
			nums: &ListNode{},
			want: []int{0},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(st *testing.T) {

			if got := tt.nums.ToArray(); !reflect.DeepEqual(got, tt.want) {
				st.Errorf("given %v, want %v", got, tt.want)
			}
		})
	}
}
