package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIgnoreRegister(t *testing.T) {
	lines := []string{"low", "UPP", "aaaBBB"}
	IgnoreRegister(&lines)
	expetcted := []string{"low", "upp", "aaabbb"}
	assert.Equal(t, lines, expetcted )
}

func TestIgnoreNFields(t *testing.T) {
	lines := []string{"line one aaaa", "second line", "b s a"}
	IgnoreNFields(&lines, 1)
	expetcted := []string{"one aaaa", "line", "s a"}
	assert.Equal(t, lines, expetcted)

	lines = []string{"line one aaaa", "second line", "b s a"}
	IgnoreNFields(&lines, 3)
	expetcted = []string{"", "", ""}
	assert.Equal(t, lines, expetcted)
}

func TestIgnoreNFieldsError(t *testing.T) {
	lines := []string{"line one aaaa", "second line", "b s a"}
	err := IgnoreNFields(&lines, -2)
	assert.NotNil(t, err)
}

func TestIgnoreNSymbols(t *testing.T) {
	lines := []string{"line one aaaa", "second line", "b s a"}
	IgnoreNSymbols(&lines, 3)
	expetcted := []string{"e one aaaa", "ond line", " a"}
	assert.Equal(t, lines, expetcted)

	lines = []string{"line one aaaa", "second line", "b s a"}
	IgnoreNSymbols(&lines, 20)
	expetcted = []string{"", "", ""}
	assert.Equal(t, lines, expetcted)
}

func TestIgnoreNSymbolsError(t *testing.T) {
	lines := []string{"line one aaaa", "second line", "b s a"}
	err := IgnoreNSymbols(&lines, -2)
	assert.NotNil(t, err)
}
