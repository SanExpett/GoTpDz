package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) { // тестим методы стека (успешные сценарии)
	stack := Create()

	stack.Push("first")
	assert.Equal(t, 1, stack.Len())

	top, _ := stack.Top()
	assert.Equal(t, "first", top)

	poped, _ := stack.Pop()
	assert.Equal(t, "first", poped)
	assert.Equal(t, 0, stack.Len())
}

func TestStackErrors(t *testing.T) {
	stack := Create()

	_, err := stack.Top() // нельзя топнуть из пустого стека
	assert.NotNil(t, err)

	_, err = stack.Pop() // нельзя попнуть из пустого стека
	assert.NotNil(t, err)
}
