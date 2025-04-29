package slice

// DiffSet 差集，已去重
func DiffSet[T comparable](src, dst []T) []T {
	srcMap := toMap[T](src)
	for _, val := range dst {
		delete(srcMap, val)
	}

	ret := make([]T, 0, len(srcMap))
	for key := range srcMap {
		ret = append(ret, key)
	}

	return ret
}

// DiffSetFunc 差集，已去重
func DiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	ret := make([]T, 0, len(src))
	for _, val := range src {
		if !ContainsFunc[T](dst, func(target T) bool {
			return equal(val, target)
		}) {
			ret = append(ret, val)
		}
	}
	return deduplicateFunc[T](ret, equal)
}
