package functional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllInt(t *testing.T) {
	// True if a given number > 10
	condition := func(i int) bool {
		return i > 10
	}

	passing := []int{12, 25, 36, 56, 90, 99}
	failing := []int{3, 5, 8, 12, 20, 32}
	var empty []int

	assert.True(t, AllInt(passing, condition))
	assert.False(t, AllInt(failing, condition))
	assert.True(t, AllInt(empty, condition))
}

func TestAllString(t *testing.T) {
	// True if a given string is longer than 2 chars
	condition := func(i string) bool {
		return len(i) > 2
	}

	passing := []string{"hello", "goodbye", "golang", "rocks"}
	failing := []string{"I", "think", "it", "doesn't"}
	var empty []string

	assert.True(t, AllString(passing, condition))
	assert.False(t, AllString(failing, condition))
	assert.True(t, AllString(empty, condition))
}

func TestAnyInt(t *testing.T) {
	// True if a given number > 10
	condition := func(i int) bool {
		return i > 10
	}

	passing := []int{12, 25, 36, 56, 90, 99}
	failing := []int{3, 5, 8, 9, 0, 2}
	var empty []int

	assert.True(t, AnyInt(passing, condition))
	assert.False(t, AnyInt(failing, condition))
	assert.False(t, AnyInt(empty, condition))
}

func TestAnyString(t *testing.T) {
	// True if a given string is longer than 2 chars
	condition := func(i string) bool {
		return len(i) > 2
	}

	passing := []string{"hello", "goodbye", "golang", "rocks"}
	failing := []string{"I", "it"}
	var empty []string

	assert.True(t, AnyString(passing, condition))
	assert.False(t, AnyString(failing, condition))
	assert.False(t, AnyString(empty, condition))
}
