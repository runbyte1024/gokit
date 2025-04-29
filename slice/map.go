package slice

// FilterMap 执行过滤以及转换
//
// 将 Src 类型的数组转换为 Dst 类型的数组， 并且根据 m 返回值进行过滤
func FilterMap[Src any, Dst any](src []Src, m func(idx int, src Src) (Dst, bool)) []Dst {
	res := make([]Dst, 0, len(src))
	for i, s := range src {
		dst, ok := m(i, s)
		if ok {
			res = append(res, dst)
		}
	}

	return res
}

// Map 执行转换
//
// 将 Src 类型的数组转换为 Dst 类型的数组
func Map[Src any, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	return FilterMap(src, func(idx int, src Src) (Dst, bool) {
		return m(idx, src), true
	})
}

func ToMap[Ele any, Key comparable](elements []Ele, fn func(element Ele) Key) map[Key]Ele {
	return ToMapV(elements, func(element Ele) (Key, Ele) {
		return fn(element), element
	})
}

func ToMapV[Ele any, Key comparable, Val any](elements []Ele, fn func(element Ele) (Key, Val)) (resultMap map[Key]Val) {
	resultMap = make(map[Key]Val, len(elements))
	for _, element := range elements {
		k, v := fn(element)
		resultMap[k] = v
	}
	return
}

func toMap[T comparable](src []T) map[T]struct{} {
	m := make(map[T]struct{}, len(src))
	for _, v := range src {
		m[v] = struct{}{}
	}
	return m
}

func deduplicateFunc[T any](data []T, equal equalFunc[T]) []T {
	newData := make([]T, 0, len(data))
	for k, v := range data {
		if !ContainsFunc[T](data[k+1:], func(target T) bool {
			return equal(target, v)
		}) {
			newData = append(newData, v)
		}
	}
	return newData
}

func deduplicate[T comparable](data []T) []T {
	dataMap := toMap[T](data)
	newData := make([]T, 0, len(dataMap))
	for key := range dataMap {
		newData = append(newData, key)
	}
	return newData
}
