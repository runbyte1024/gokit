package slice

import (
	"fmt"
	"github.com/runbyte1024/gokit/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		addVal    int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "索引 0",
			slice:     []int{123, 100},
			addVal:    233,
			index:     0,
			wantSlice: []int{233, 123, 100},
		},
		{
			name:    "索引 -1",
			slice:   []int{123, 100},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfBounds(2, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addVal, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func ExampleAdd() {
	res, _ := Add[int]([]int{1, 2, 3, 4}, 5, 2)
	fmt.Println(res)
	_, err := Add[int]([]int{1, 2, 3, 4}, 5, -1)
	fmt.Println(err)
	// Output:
	// [1 2 5 3 4]
	// gokit: 索引越界, 长度: 4, 索引: -1
}
