package slice

// Find 查找元素
//
// 若未找到返回对应类型的零值以及false
func Find[T any](src []T, match matchFunc[T]) (T, bool) {
	for _, v := range src {
		if match(v) {
			return v, true
		}
	}
	var t T
	return t, false
}

// FindAll 查找符合条件的所有元素
func FindAll[T any](src []T, match matchFunc[T]) []T {
	res := make([]T, 0, len(src)>>3+1)
	for _, v := range src {
		if match(v) {
			res = append(res, v)
		}
	}
	return res
}
