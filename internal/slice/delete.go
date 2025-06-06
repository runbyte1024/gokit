package slice

import (
	"github.com/runbyte1024/gokit/internal/errs"
)

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfBounds(length, index)
	}

	res := src[index]

	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	src = src[:length-1]
	return src, res, nil
}
