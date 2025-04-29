package slice

import (
	"github.com/runbyte1024/gokit/internal/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantVal   int
		wantErr   error
	}{
		{
			name:      "索引 0",
			slice:     []int{1, 2},
			index:     0,
			wantSlice: []int{2},
			wantVal:   1,
		},
		{
			name:      "索引中间",
			slice:     []int{1, 2, 3},
			index:     1,
			wantSlice: []int{1, 3},
			wantVal:   2,
		},
		{
			name:    "索引越界",
			slice:   []int{1, 2, 3},
			index:   10,
			wantErr: errs.NewErrIndexOutOfBounds(3, 10),
		},
		{
			name:    "索引小于0",
			slice:   []int{1, 2, 3},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfBounds(3, -1),
		},
		{
			name:      "索引尾部",
			slice:     []int{1, 2, 3},
			index:     2,
			wantSlice: []int{1, 2},
			wantVal:   3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.slice, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}
