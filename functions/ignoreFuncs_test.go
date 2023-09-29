package functions_test

import (
	"testing"
	
	"github.com/SanExpett/GoDz1P1/functions"
	"github.com/stretchr/testify/assert"
)

func TestIgnoreRegister(t *testing.T) {
	t.Parallel()

	lines := []string{"low", "UPP", "aaaBBB"}
	functions.IgnoreRegister(&lines)
	
	expetcted := []string{"low", "upp", "aaabbb"}
	assert.Equal(t, lines, expetcted )
}

func TestIgnoreNFields(t *testing.T) {
	t.Parallel()

	lines := []string{"line one aaaa", "second line", "b s a"}
	_ = functions.IgnoreNFields(&lines, 1)

	expetcted := []string{"one aaaa", "line", "s a"}
	assert.Equal(t, lines, expetcted)

	lines = []string{"line one aaaa", "second line", "b s a"}
	_ = functions.IgnoreNFields(&lines, 3)

	expetcted = []string{"", "", ""}
	assert.Equal(t, lines, expetcted)
}

func TestIgnoreNFieldsError(t *testing.T) {
	t.Parallel()

	lines := []string{"line one aaaa", "second line", "b s a"}
	err := functions.IgnoreNFields(&lines, -2)
	assert.NotNil(t, err)
}

func TestIgnoreNSymbols(t *testing.T) {
	t.Parallel()

	lines := []string{"line one aaaa", "second line", "b s a"}
	_ = functions.IgnoreNSymbols(&lines, 3)

	expetcted := []string{"e one aaaa", "ond line", " a"}
	assert.Equal(t, lines, expetcted)

	lines = []string{"line one aaaa", "second line", "b s a"}
	_ = functions.IgnoreNSymbols(&lines, 20)

	expetcted = []string{"", "", ""}
	assert.Equal(t, lines, expetcted)
}

func TestIgnoreNSymbolsError(t *testing.T) {
	t.Parallel()
	
	lines := []string{"line one aaaa", "second line", "b s a"}
	err := functions.IgnoreNSymbols(&lines, -2)
	assert.NotNil(t, err)
}
