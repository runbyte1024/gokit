package slice

// UnionSet 并集，已去重
func UnionSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap[T](src), toMap[T](dst)

	for k := range srcMap {
		dstMap[k] = struct{}{}
	}

	ret := make([]T, 0, len(dstMap))
	for k := range dstMap {
		ret = append(ret, k)
	}
	return ret
}

// UnionSetFunc 并集，已去重
func UnionSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	ret := make([]T, 0, len(src)+len(dst))

	ret = append(ret, src...)
	ret = append(ret, dst...)
	return deduplicateFunc[T](ret, equal)
}
