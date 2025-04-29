package slice

import (
	"github.com/runbyte1024/gokit/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		addValue  int
		index     int
		wantSlice []int
		wantError error
	}{
		{
			name:      "头部插入",
			slice:     []int{2, 3},
			addValue:  1,
			index:     0,
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "中间插入",
			slice:     []int{1, 3},
			addValue:  2,
			index:     1,
			wantSlice: []int{1, 2, 3},
		},
		{
			name:      "索引尾部",
			slice:     []int{1, 2, 3},
			addValue:  4,
			index:     2,
			wantSlice: []int{1, 2, 4, 3},
		},
		{
			name:      "尾部添加",
			slice:     []int{1, 2, 3},
			addValue:  4,
			index:     3,
			wantSlice: []int{1, 2, 3, 4},
		},
		{
			name:      "索引超出",
			slice:     []int{1, 2, 3},
			index:     10,
			wantError: errs.NewErrIndexOutOfBounds(3, 10),
		},
		{
			name:      "索引小于",
			slice:     []int{1, 2, 3},
			index:     -1,
			wantError: errs.NewErrIndexOutOfBounds(3, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addValue, tc.index)
			assert.Equal(t, tc.wantError, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
