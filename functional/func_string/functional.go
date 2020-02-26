package func_string

func All(collection []string, fn func(string) bool) bool {
	for i := 0; i < len(collection); i++ {
		if !fn(collection[i]) {
			return false
		}
	}
	return true
}

func Any(collection []string, fn func(string) bool) bool {
	for i := 0; i < len(collection); i++ {
		if fn(collection[i]) {
			return true
		}
	}
	return false
}

func Filter(collection []string, fn func(string) bool) []string {
	var mapped = make([]string, len(collection))

	filtered := 0
	for i := 0; i < len(collection); i++ {
		if fn(collection[i]) {
			mapped[filtered] = collection[i]
			filtered++
		}
	}

	return mapped[:filtered]
}

func FilterNot(collection []string, fn func(string) bool) []string {
	var mapped = make([]string, len(collection))

	filtered := 0
	for i := 0; i < len(collection); i++ {
		if !fn(collection[i]) {
			mapped[filtered] = collection[i]
			filtered++
		}
	}

	return mapped[:filtered]
}

func Flatten(collection [][]string) []string {
	slots := 0
	for i := 0; i < len(collection); i++ {
		slots += len(collection[i])
	}

	flattened := make([]string, slots)

	for i, startingIndex := 0, 0; i < len(collection); i++ {
		sizeOfCurrent := len(collection[i])
		for j := 0; j < sizeOfCurrent; j++ {
			flattened[startingIndex+j] = collection[i][j]
		}
		startingIndex += sizeOfCurrent
	}

	return flattened
}

func Fold(collection []string, initial string, fn func(string, string) string) (result string) {
	result = initial
	for i := 0; i < len(collection); i++ {
		result = fn(result, collection[i])
	}
	return
}

func FoldR(collection []string, initial string, fn func(string, string) string) (result string) {
	result = initial
	for i := len(collection) - 1; i >= 0; i-- {
		result = fn(result, collection[i])
	}
	return
}

func ForEach(collection []string, fn func(string)) {
	for i := 0; i < len(collection); i++ {
		fn(collection[i])
	}
}

func IndexOf(collection []string, val string) int {
	for i := 0; i < len(collection); i++ {
		if collection[i] == val {
			return i
		}
	}
	return -1
}

func Map(collection []string, fn func(string) string) []string {
	size := len(collection)
	mapped := make([]string, size)
	for i := 0; i < size; i++ {
		mapped[i] = fn(collection[i])
	}
	return mapped
}

func Reduce(collection []string, fn func(string, string) string) (result string) {
	size := len(collection)
	if size == 0 {
		return
	}
	result = collection[0]
	for i := 1; i < size; i++ {
		result = fn(result, collection[i])
	}
	return
}

func ReduceR(collection []string, fn func(string, string) string) (result string) {
	size := len(collection)
	if size == 0 {
		return
	}
	result = collection[size-1]
	for i := size - 2; i >= 0; i-- {
		result = fn(result, collection[i])
	}
	return
}
