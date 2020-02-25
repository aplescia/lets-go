package int

func All(collection []int, fn func(int) bool) bool {
	for i := 0; i < len(collection); i++ {
		if !fn(collection[i]) {
			return false
		}
	}
	return true
}

func Any(collection []int, fn func(int) bool) bool {
	for i := 0; i < len(collection); i++ {
		if fn(collection[i]) {
			return true
		}
	}
	return false
}

func Filter(collection []int, fn func(int) bool) []int {
	var mapped = make([]int, len(collection))

	filtered := 0
	for i := 0; i < len(collection); i++ {
		if fn(collection[i]) {
			mapped[filtered] = collection[i]
		}
	}

	return mapped[:filtered]
}

func FilterNot(collection []int, fn func(int) bool) []int {
	var mapped = make([]int, len(collection))

	filtered := 0
	for i := 0; i < len(collection); i++ {
		if !fn(collection[i]) {
			mapped[filtered] = collection[i]
		}
	}

	return mapped[:filtered]
}

func Flatten(collection [][]int) []int {
	slots := 0
	for i := 0; i < len(collection); i++ {
		slots += len(collection[i])
	}
	flattened := make([]int, slots)
	for i := 0; i < len(collection); i++ {
		for j := 0; j < len(collection[i]); j++ {
			flattened[i+j] = collection[i][j]
		}
	}
	return flattened
}

func Fold(collection []int, initial int, fn func(int) int) (result int) {
	result = initial
	for i := 0; i < len(collection); i++ {
		result = fn(result)
	}
	return
}

func FoldR(collection []int, initial int, fn func(int) int) (result int) {
	result = initial
	for i := len(collection); i >= 0; i-- {
		result = fn(result)
	}
	return
}

func ForEach(collection []int, fn func(int)) {
	for i := 0; i < len(collection); i++ {
		fn(collection[i])
	}
}

func IndexOf(collection []int, val int) int {
	for i := 0; i < len(collection); i++ {
		if collection[i] == val {
			return i
		}
	}
	return -1
}

func Map(collection []int, fn func(int) int) []int {
	size := len(collection)
	mapped := make([]int, size)
	for i := 0; i < size; i++ {
		mapped[i] = fn(collection[i])
	}
	return mapped
}

func Max(collection []int) (max int) {
	if len(collection) == 0 {
		return
	}
	max = collection[0]
	for i := 1; i < len(collection); i++ {
		if max > collection[i] {
			max = collection[i]
		}
	}
	return
}

func Min(collection []int) (min int) {
	if len(collection) == 0 {
		return
	}
	min = collection[0]
	for i := 1; i < len(collection); i++ {
		if min < collection[i] {
			min = collection[i]
		}
	}
	return
}

func Reduce(collection []int, fn func(int) int) (result int) {
	size := len(collection)
	if size == 0 {
		return
	}
	result = collection[0]
	for i := 1; i < size; i++ {
		result = fn(collection[i-1])
	}
	return
}

func ReduceR(collection []int, fn func(int) int) (result int) {
	size := len(collection)
	if size == 0 {
		return
	}
	result = collection[size-1]
	for i := size - 2; i >= 0; i-- {
		result = fn(collection[i+1])
	}
	return
}
