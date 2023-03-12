package roman_to_integer

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Solve(t *testing.T) {
	s := New()

	tests := []struct {
		input string
		want  int
	}{
		{
			input: "III",
			want:  3,
		},
		{
			input: "LVIII",
			want:  58,
		},
		{
			input: "MCMXCIV",
			want:  1994,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(st *testing.T) {
			if got := s.Solve(tt.input); !reflect.DeepEqual(got, tt.want) {
				st.Errorf("given %v, want %v", got, tt.want)
			}
		})
	}
}
