package slice

import "github.com/runbyte1024/gokit/internal/slice"

// Add 在 index 处添加 element，index 范围应在 [0, len(src)]，若 index == len(src) 则在末尾添加 element
func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	return slice.Add[Src](src, element, index)
}
