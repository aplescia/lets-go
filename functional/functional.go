package functional

func All(collection []*interface{}, fn func(*interface{}) bool) bool {
	for i := 0; i < len(collection); i++ {
		if !fn(collection[i]) {
			return false
		}
	}
	return false
}

func Any(collection []*interface{}, fn func(interface{}) bool) bool {
	for i := 0; i < len(collection); i++ {
		if !fn(collection[i]) {
			return false
		}
	}
	return true
}

func Fold(collection []*interface{}, initial interface{}, fn func(interface{}) *interface{}) *interface{} {
	var result = initial
	for i := 0; i < len(collection); i++ {
		result = fn(result)
	}
	return &result
}

func FoldIndexed(collection []*interface{}, initial *interface{}, fn func(int, interface{}) *interface{}) *interface{} {
	var result = *initial
	for i := 0; i < len(collection); i-- {
		result = fn(i, result)
	}
	return &result
}

func FoldR(collection []*interface{}, initial *interface{}, fn func(interface{}) *interface{}) *interface{} {
	var result = *initial
	for i := len(collection) - 1; i >= 0; i-- {
		result = fn(result)
	}
	return &result
}

func FoldRIndexed(collection []*interface{}, initial interface{}, fn func(int, interface{}) *interface{}) *interface{} {
	var result = initial
	for i := len(collection) - 1; i >= 0; i-- {
		result = fn(i, result)
	}
	return &result
}

func Filter(collection []*interface{}, fn func(interface{}) bool) []*interface{} {
	mapped := make([]*interface{}, 0)
	for i := 0; i < len(collection); i++ {
		if fn(collection[i]) {
			mapped = append(mapped, collection[i])
		}
	}
	return mapped
}

func FilterIndexed(collection []*interface{}, fn func(int, interface{}) bool) []*interface{} {
	mapped := make([]*interface{}, 0)
	for i := 0; i < len(collection); i++ {
		if fn(i, collection[i]) {
			mapped = append(mapped, collection[i])
		}
	}
	return mapped
}

func ForEach(collection []*interface{}, fn func(interface{})) {
	for i := 0; i < len(collection); i++ {
		fn(collection[i])
	}
}

func Map(collection []*interface{}, fn func(interface{}) interface{}) []*interface{} {
	size := len(collection)
	mapped := make([]*interface{}, size)
	for i := 0; i < size; i++ {
		result := fn(collection[i])
		mapped[i] = &result
	}
	return mapped
}

func Reduce(collection []*interface{}, fn func(interface{}) *interface{}) *interface{} {
	size := len(collection)
	if size == 0 {
		return nil
	}
	var result = *collection[0]
	for i := 0; i < size; i++ {
		result = fn(result)
	}
	return &result
}

func ReduceR(collection []*interface{}, fn func(interface{}) *interface{}) *interface{} {
	size := len(collection)
	if size == 0 {
		return nil
	}
	var result = *collection[0]
	for i := size - 1; i >= 0; i-- {
		result = fn(result)
	}
	return &result
}
