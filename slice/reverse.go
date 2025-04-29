package slice

// Reverse 返回一个反转后的新切片
func Reverse[T any](src []T) []T {
	ret := make([]T, 0, len(src))
	for i := len(src) - 1; i >= 0; i-- {
		ret = append(ret, src[i])
	}
	return ret
}

// ReverseSelf 反转切片
func ReverseSelf[T any](src []T) {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
}
