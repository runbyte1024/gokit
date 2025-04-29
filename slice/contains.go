package slice

// Contains 判断 src 是否存在 target
func Contains[T comparable](src []T, target T) bool {
	return ContainsFunc[T](src, func(t T) bool {
		return t == target
	})
}

// ContainsFunc 判断 src 是否存在 target
func ContainsFunc[T any](src []T, equal func(target T) bool) bool {
	for _, v := range src {
		if equal(v) {
			return true
		}
	}
	return false
}

// ContainsAny 判断 src 里面是否存在 dst 中的任何一个元素
func ContainsAny[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; exist {
			return true
		}
	}
	return false
}

// ContainsAnyFunc 判断 src 里面是否存在 dst 中的任何一个元素
func ContainsAnyFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, v := range src {
		for _, t := range dst {
			if equal(v, t) {
				return true
			}
		}
	}
	return false
}

// ContainsAll 判断 src 里面是否存在 dst 中的所有元素
func ContainsAll[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; !exist {
			return false
		}
	}
	return true
}

// ContainsAllFunc 判断 src 里面是否存在 dst 中的所有元素
func ContainsAllFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, val := range dst {
		if !ContainsFunc[T](src, func(target T) bool {
			return equal(val, target)
		}) {
			return false
		}
	}
	return true
}
