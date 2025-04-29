package slice

import "github.com/runbyte1024/gokit/internal/errs"

func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)

	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfBounds(length, index)
	}

	var zeroValue T
	src = append(src, zeroValue)
	for i := len(src) - 1; i > index; i-- {
		src[i] = src[i-1]
	}

	src[index] = element
	return src, nil
}
