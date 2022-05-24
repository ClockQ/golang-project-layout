package channel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildLazyIntEvaluator(t *testing.T) {
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}
	even := BuildLazyIntEvaluator(evenFunc, 0)

	fmt.Println(even())
	assert.Equal(t, 2, even(), "BuildLazyIntEvaluator failed")
	fmt.Println(even())
	assert.Equal(t, 6, even(), "BuildLazyIntEvaluator failed")
	fmt.Println(even())
	assert.Equal(t, 10, even(), "BuildLazyIntEvaluator failed")
	fmt.Println(even())
}
