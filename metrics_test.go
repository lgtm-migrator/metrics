package metrics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func func1() string {
	return func2()
}

func func2() string {
	return func3()
}

func func3() string {
	return CallerFuncName(2)
}

func TestCallerFuncName(t *testing.T) {
	func1Name := func1()
	assert.Equal(t, func1Name, "func1")
}

func TestGetFuncName(t *testing.T) {
	func1Name := GetFuncName(func1)
	assert.Equal(t, func1Name, "func1")
}
