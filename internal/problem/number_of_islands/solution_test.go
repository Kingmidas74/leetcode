package number_of_islands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solve(t *testing.T) {
	s := New()

	tests := []struct {
		input    []string
		rowCount int
		colCount int
		want     int
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			input:    []string{"1", "1", "1", "1", "0", "1", "1", "0", "1", "0", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0"},
			rowCount: 4,
			colCount: 5,
			want:     1,
			wantErr:  assert.NoError,
		},
		{
			input:    []string{"1", "1", "0", "0", "0", "1", "1", "0", "0", "0", "0", "0", "1", "0", "0", "0", "0", "0", "1", "1"},
			rowCount: 4,
			colCount: 5,
			want:     3,
			wantErr:  assert.NoError,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(st *testing.T) {
			got, err := s.Solve(tt.input, tt.rowCount, tt.colCount)
			if !tt.wantErr(st, err) {
				t.Errorf("err: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				st.Errorf("given %v, want %v", got, tt.want)
			}
		})
	}
}
