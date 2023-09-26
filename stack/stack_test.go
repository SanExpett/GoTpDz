package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) { // тестим методы стека (успешные сценарии)
	t.Parallel()

	stack := Create()
	assert.Equal(t, 0, stack.Len())

	stack.Push("first")
	assert.Equal(t, 1, stack.Len())

	top, _ := stack.Top()
	assert.Equal(t, "first", top)
	assert.Equal(t, 1, stack.Len())

	poped, _ := stack.Pop()
	assert.Equal(t, "first", poped)
	assert.Equal(t, 0, stack.Len())

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	stack.Push("4")
	assert.Equal(t, 4, stack.Len())
	poped, _ = stack.Pop()
	assert.Equal(t, "4", poped)
	assert.Equal(t, 3, stack.Len())
}

func TestStackErrors(t *testing.T) {
	t.Parallel()
	
	stack := Create()

	_, err := stack.Top() // нельзя топнуть из пустого стека
	assert.NotNil(t, err)

	_, err = stack.Pop() // нельзя попнуть из пустого стека
	assert.NotNil(t, err)
}
