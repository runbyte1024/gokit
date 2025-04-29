package slice

// SymmetricDiffSet 对称差集，已去重
func SymmetricDiffSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap[T](src), toMap[T](dst)

	for k := range dstMap {
		if _, ok := srcMap[k]; ok {
			delete(srcMap, k)
		} else {
			srcMap[k] = struct{}{}
		}
	}

	ret := make([]T, 0, len(srcMap))
	for k := range srcMap {
		ret = append(ret, k)
	}

	return ret
}

// SymmetricDiffSetFunc 对称差集，已去重
func SymmetricDiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var ret []T

	for _, v := range src {
		if !ContainsFunc[T](dst, func(target T) bool {
			return equal(v, target)
		}) {
			ret = append(ret, v)
		}
	}

	for _, v := range dst {
		if !ContainsFunc[T](src, func(target T) bool {
			return equal(v, target)
		}) {
			ret = append(ret, v)
		}
	}

	return deduplicateFunc(ret, equal)
}
