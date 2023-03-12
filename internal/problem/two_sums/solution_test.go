package two_sums

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
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
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
