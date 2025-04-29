package slice

func calCapacity(c, l int) (int, bool) {
	if c <= 64 {
		return c, false
	}

	if c > 2048 && c/l >= 2 {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}

	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}

	return c, false
}

func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	capacity, b := calCapacity(c, l)
	if !b {
		return src
	}
	s := make([]T, 0, capacity)
	s = append(s, src...)
	return s
}
