package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) { // тестим методы стека (успешные сценарии)
	st := Create() 

	st.Push("first")
	assert.Equal(t, 1, st.Len())

	top, _ := st.Top()
	assert.Equal(t, "first", top)

	poped, _ := st.Pop()
	assert.Equal(t, "first", poped)
	assert.Equal(t, 0, st.Len())
}

func TestStackErrors(t *testing.T) {
	st := Create() 

	_, err := st.Top() // нельзя топнуть из пустого стека
	assert.NotNil(t, err)

	_, err = st.Pop() // нельзя попнуть из пустого стека
	assert.NotNil(t, err)
}
