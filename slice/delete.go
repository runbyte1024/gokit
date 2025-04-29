package slice

import "github.com/runbyte1024/gokit/internal/slice"

// Delete 删除指定 index 处的元素
func Delete[Src any](src []Src, index int) ([]Src, error) {
	res, _, err := slice.Delete[Src](src, index)
	return res, err
}

// FilterDelete 删除符合条件的元素
func FilterDelete[Src any](src []Src, m func(idx int, src Src) bool) []Src {
	emptyPos := 0
	for idx := range src {
		if m(idx, src[idx]) {
			continue
		}
		src[emptyPos] = src[idx]
		emptyPos++
	}
	return src[:emptyPos]
}
