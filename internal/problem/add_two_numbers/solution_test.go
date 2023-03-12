package add_two_numbers

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Solve(t *testing.T) {
	s := New()

	tests := []struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			arr1: []int{2, 4, 3},
			arr2: []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			arr1: []int{0},
			arr2: []int{0},
			want: []int{0},
		},
		{
			arr1: []int{9, 9, 9, 9, 9, 9, 9},
			arr2: []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(st *testing.T) {
			if got := s.Solve(tt.arr1, tt.arr2); !reflect.DeepEqual(got, tt.want) {
				st.Errorf("given %v, want %v", got, tt.want)
			}
		})
	}
}
