package slice

import "github.com/runbyte1024/gokit"

// Max 返回最大值
func Max[T gokit.RealNumber](ts []T) T {
	res := ts[0]
	for i := 1; i < len(ts); i++ {
		if ts[i] > res {
			res = ts[i]
		}
	}
	return res
}

// Min 返回最小值
func Min[T gokit.RealNumber](ts []T) T {
	res := ts[0]
	for i := 1; i < len(ts); i++ {
		if ts[i] < res {
			res = ts[i]
		}
	}
	return res
}

// Sum 求和
func Sum[T gokit.Number](ts []T) T {
	var res T
	for _, n := range ts {
		res += n
	}
	return res
}
