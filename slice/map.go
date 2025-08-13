package slice

// Map 转换切片类型
// []A --> []B
func Map[Src any, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	for i, s := range src {
		dst[i] = m(i, s)
	}
	return dst
}

// ToMap 将[]Ele映射到map[Key]Ele
// 从Ele中提取Key的函数fn由使用者提供
// 相同的key会覆盖
// 即使传入的字符串为nil，也保证返回的map是一个空map而不是nil
func ToMap[Ele any, Key comparable](elements []Ele, fn func(element Ele) Key) map[Key]Ele {
	return ToMapV(
		elements,
		func(element Ele) (Key, Ele) {
			return fn(element), element
		})
}

// ToMapV 将[]Ele映射到map[Key]Val
// 从Ele中提取Key和Val的函数fn由使用者提供
// 相同的key会覆盖
// 即使传入的字符串为nil，也保证返回的map是一个空map而不是nil
func ToMapV[Ele any, Key comparable, Val any](elements []Ele, fn func(element Ele) (Key, Val)) (resultMap map[Key]Val) {
	resultMap = make(map[Key]Val, len(elements))
	for _, element := range elements {
		k, v := fn(element)
		resultMap[k] = v
	}
	return
}
