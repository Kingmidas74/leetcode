package maximum_subarray

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Solve(t *testing.T) {
	s := New()

	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			nums: []int{1},
			want: 1,
		},
		{
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(st *testing.T) {
			if got := s.Solve(tt.nums); !reflect.DeepEqual(got, tt.want) {
				st.Errorf("given %v, want %v", got, tt.want)
			}
		})
	}
}
