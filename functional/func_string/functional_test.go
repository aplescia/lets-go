package func_string_test

import (
	"fmt"
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

	assert.True(t, All(passing, condition))
	assert.False(t, All(failing, condition))
	assert.True(t, All(empty, condition))
}

func TestAny(t *testing.T) {
	condition := func(i string) bool {
		return len(i) > 7
	}

	passing := []string{"Perfectly", "Balanced"}
	failing := []string{"As", "All", "Things", "Should", "be"}

	assert.True(t, Any(passing, condition))
	assert.False(t, Any(failing, condition))
	assert.False(t, Any(empty, condition))
}

func TestFilter(t *testing.T) {
	condition := func(i string) bool {
		return len(i) >= 5
	}

	initial := []string{"As", "All", "Things", "Should", "be"}

	filtered := Filter(initial, condition)

	assert.Equal(t, 2, len(filtered))
	assert.EqualValues(t, filtered, []string{"Things", "Should"})

	assert.Equal(t,0, len(Filter(empty, condition)))
}

func TestFilterNot(t *testing.T) {
	condition := func(i string) bool {
		return len(i) >= 5
	}

	initial := []string{"As", "All", "Things", "Should", "be"}

	filtered := FilterNot(initial, condition)

	assert.Equal(t, 3, len(filtered))
	assert.EqualValues(t, []string{"As", "All", "be"}, filtered)

	assert.Equal(t,0, len(FilterNot(empty, condition)))
}

func TestFlatten(t *testing.T) {
	slice1 := []string{"Perfectly", "Balanced"}
	slice2 := []string{"As", "All", "Things", "Should", "be"}
	slice3 := []string{".", "I", "AM", "INEVITABLE"}
	matrix := [][]string{slice1, slice2, slice3}

	assert.EqualValues(t, []string{"Perfectly", "Balanced", "As", "All", "Things", "Should", "be", ".", "I", "AM", "INEVITABLE"}, Flatten(matrix))
	assert.Equal(t, len(empty), len(Flatten([][]string{empty})))
}

func TestFold(t *testing.T) {
	foldFn := func(x string, y string) string {
		return x + " " + y
	}

	sliceToFold := []string{"As", "All", "Things", "Should", "Be"}

	assert.Equal(t, Fold(sliceToFold, "Perfectly Balanced", foldFn), "Perfectly Balanced As All Things Should Be")
	assert.Equal(t, Fold(empty, "", foldFn), "")
}

func TestFoldR(t *testing.T) {
	foldFn := func(x string, y string) string {
		return x + " " + y
	}

	sliceToFold := []string{"As", "All", "Things", "Should", "Be"}

	assert.Equal(t,"We Be Should Things All As", FoldR(sliceToFold, "We", foldFn))
	assert.Equal(t, "", FoldR(empty, "", foldFn))
}

func TestForEach(t *testing.T) {
	sliceToPrint := []string{"As", "All", "Things", "Should", "Be"}

	ForEach(sliceToPrint, func(x string){fmt.Println(x)})
}

func TestIndexOf(t *testing.T) {
	sliceToFindIndex := []string{"As", "All", "Things", "Should", "Be"}

	assert.Equal(t, 2, IndexOf(sliceToFindIndex, "Things"))
	assert.Equal(t, -1, IndexOf(empty, "Inevitable"))
}

func TestMap(t *testing.T) {
	mapFn := func(x string) string {
		return x + "<->"
	}

	slice := []string{"As", "All", "Things", "Should", "Be"}

	assert.EqualValues(t, []string{"As<->", "All<->", "Things<->", "Should<->", "Be<->"}, Map(slice, mapFn))
	assert.EqualValues(t, []string{}, Map(empty, mapFn))
}

func TestReduce(t *testing.T) {
	foldFn := func(x string, y string) string {
		return x + " " + y
	}

	sliceToFold := []string{"As", "All", "Things", "Should", "Be"}

	assert.Equal(t, Reduce(sliceToFold,  foldFn), "As All Things Should Be")
	assert.Equal(t, Reduce(empty, foldFn), "")

}

func TestReduceR(t *testing.T) {
	foldFn := func(x string, y string) string {
		return x + " " + y
	}

	sliceToFold := []string{"As", "All", "Things", "Should", "Be"}

	assert.Equal(t, ReduceR(sliceToFold, foldFn), "Be Should Things All As")
	assert.Equal(t, ReduceR(empty,  foldFn), "")
}
