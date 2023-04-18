package three_sums

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Solve(t *testing.T) {
	s := New()

	tests := []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{
			nums:   []int{-1, 0, 1, 2, -1, -4},
			target: 0,
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			nums:   []int{0, 1, 1},
			target: 0,
			want:   [][]int{},
		},
		{
			nums:   []int{0, 0, 0},
			target: 0,
			want: [][]int{
				{0, 0, 0},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(st *testing.T) {
			if got := s.Solve(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				st.Errorf("given %v, want %v", got, tt.want)
			}
		})
	}
}
