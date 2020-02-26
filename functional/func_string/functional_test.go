package func_string_test

import (
	"fmt"
	"github.com/Chewy-Inc/lets-go/functional/func_string"
	"github.com/stretchr/testify/assert"
	"testing"
)

var empty []string

func TestAll(t *testing.T) {
	condition := func(i string) bool {
		return len(i) >= 5
	}

	passing := []string{"Perfectly", "Balanced"}
	failing := []string{"As", "All", "Things", "Should", "be"}

	assert.True(t, func_string.All(passing, condition))
	assert.False(t, func_string.All(failing, condition))
	assert.True(t, func_string.All(empty, condition))
}

func TestAny(t *testing.T) {
	condition := func(i string) bool {
		return len(i) > 7
	}

	passing := []string{"Perfectly", "Balanced"}
	failing := []string{"As", "All", "Things", "Should", "be"}

	assert.True(t, func_string.Any(passing, condition))
	assert.False(t, func_string.Any(failing, condition))
	assert.False(t, func_string.Any(empty, condition))
}

func TestFilter(t *testing.T) {
	condition := func(i string) bool {
		return len(i) >= 5
	}

	initial := []string{"As", "All", "Things", "Should", "be"}

	filtered := func_string.Filter(initial, condition)

	assert.Equal(t, 2, len(filtered))
	assert.EqualValues(t, filtered, []string{"Things", "Should"})

	assert.Equal(t,0, len(func_string.Filter(empty, condition)))
}

func TestFilterNot(t *testing.T) {
	condition := func(i string) bool {
		return len(i) >= 5
	}

	initial := []string{"As", "All", "Things", "Should", "be"}

	filtered := func_string.FilterNot(initial, condition)

	assert.Equal(t, 3, len(filtered))
	assert.EqualValues(t, []string{"As", "All", "be"}, filtered)

	assert.Equal(t,0, len(func_string.FilterNot(empty, condition)))
}

func TestFlatten(t *testing.T) {
	slice1 := []string{"Perfectly", "Balanced"}
	slice2 := []string{"As", "All", "Things", "Should", "be"}
	slice3 := []string{".", "I", "AM", "INEVITABLE"}
	matrix := [][]string{slice1, slice2, slice3}

	assert.EqualValues(t, []string{"Perfectly", "Balanced", "As", "All", "Things", "Should", "be", ".", "I", "AM", "INEVITABLE"}, func_string.Flatten(matrix))
	assert.Equal(t, len(empty), len(func_string.Flatten([][]string{empty})))
}

func TestFold(t *testing.T) {
	foldFn := func(x string, y string) string {
		return x + " " + y
	}

	sliceToFold := []string{"As", "All", "Things", "Should", "Be"}

	assert.Equal(t, func_string.Fold(sliceToFold, "Perfectly Balanced", foldFn), "Perfectly Balanced As All Things Should Be")
	assert.Equal(t, func_string.Fold(empty, "", foldFn), "")
}

func TestFoldR(t *testing.T) {
	foldFn := func(x string, y string) string {
		return x + " " + y
	}

	sliceToFold := []string{"As", "All", "Things", "Should", "Be"}

	assert.Equal(t,"We Be Should Things All As", func_string.FoldR(sliceToFold, "We", foldFn))
	assert.Equal(t, "", func_string.FoldR(empty, "", foldFn))
}

func TestForEach(t *testing.T) {
	sliceToPrint := []string{"As", "All", "Things", "Should", "Be"}

	func_string.ForEach(sliceToPrint, func(x string){fmt.Println(x)})
}

func TestIndexOf(t *testing.T) {
	sliceToFindIndex := []string{"As", "All", "Things", "Should", "Be"}

	assert.Equal(t, 2, func_string.IndexOf(sliceToFindIndex, "Things"))
	assert.Equal(t, -1, func_string.IndexOf(empty, "Inevitable"))
}

func TestMap(t *testing.T) {
	mapFn := func(x string) string {
		return x + "<->"
	}

	slice := []string{"As", "All", "Things", "Should", "Be"}

	assert.EqualValues(t, []string{"As<->", "All<->", "Things<->", "Should<->", "Be<->"}, func_string.Map(slice, mapFn))
	assert.EqualValues(t, []string{}, func_string.Map(empty, mapFn))
}

func TestReduce(t *testing.T) {
	foldFn := func(x string, y string) string {
		return x + " " + y
	}

	sliceToFold := []string{"As", "All", "Things", "Should", "Be"}

	assert.Equal(t, func_string.Reduce(sliceToFold,  foldFn), "As All Things Should Be")
	assert.Equal(t, func_string.Reduce(empty, foldFn), "")

}

func TestReduceR(t *testing.T) {
	foldFn := func(x string, y string) string {
		return x + " " + y
	}

	sliceToFold := []string{"As", "All", "Things", "Should", "Be"}

	assert.Equal(t, func_string.ReduceR(sliceToFold, foldFn), "Be Should Things All As")
	assert.Equal(t, func_string.ReduceR(empty,  foldFn), "")
}
