package slice

// Index 返回指定元素的索引
//
// -1 表示未找到
func Index[T comparable](src []T, dst T) int {
	return IndexFunc(src, func(src T) bool {
		return dst == src
	})
}

// IndexFunc 返回条件匹配的第一个元素的索引
//
// -1 表示未找到
func IndexFunc[T any](src []T, match matchFunc[T]) int {
	for idx, v := range src {
		if match(v) {
			return idx
		}
	}
	return -1
}

// LastIndex 返回指定元素的最后一个索引
func LastIndex[T comparable](src []T, dst T) int {
	return LastIndexFunc(src, func(src T) bool {
		return dst == src
	})
}

// LastIndexFunc 返回条件匹配的最后一个元素的索引
//
// -1 表示未找到
func LastIndexFunc[T any](src []T, match matchFunc[T]) int {
	for i := len(src) - 1; i >= 0; i-- {
		if match(src[i]) {
			return i
		}
	}

	return -1
}

// IndexAll 返回和指定元素相等的所有元素的索引
func IndexAll[T comparable](src []T, dst T) []int {
	return IndexAllFunc(src, func(src T) bool {
		return dst == src
	})
}

// IndexAllFunc 返回条件匹配的所有元素的索引
func IndexAllFunc[T any](src []T, match matchFunc[T]) []int {
	indexes := make([]int, 0, len(src))
	for idx, v := range src {
		if match(v) {
			indexes = append(indexes, idx)
		}
	}
	return indexes
}
