package slice

// IntersectSet 交集，已去重
func IntersectSet[T comparable](src []T, dst []T) []T {
	srcMap := toMap[T](src)
	ret := make([]T, 0, len(src))
	for _, v := range dst {
		if _, exist := srcMap[v]; exist {
			ret = append(ret, v)
		}
	}
	return deduplicate[T](ret)
}

// IntersectSetFunc 交集，已去重
func IntersectSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	ret := make([]T, 0, len(src))
	for _, v := range dst {
		if ContainsFunc[T](src, func(target T) bool {
			return equal(target, v)
		}) {
			ret = append(ret, v)
		}
	}
	return deduplicateFunc[T](ret, equal)
}
