package calculate_test

import (
	"testing"

	"github.com/SanExpett/TpGoDz/calculate"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	t.Parallel()

	input := "(1+2)-3"
	actual, _ := calculate.Calculate(input)
	assert.Equal(t, "0", actual)

	input = "(1+2)*3"
	actual, _ = calculate.Calculate(input)
	assert.Equal(t, "9", actual)

	input = "(14+16-3)/(3*3)"
	actual, _ = calculate.Calculate(input)
	assert.Equal(t, "3", actual)

	input = "10*-1"
	actual, _ = calculate.Calculate(input)
	assert.Equal(t, "-10", actual)

	input = "-3+10"
	actual, _ = calculate.Calculate(input)
	assert.Equal(t, "7", actual)

	input = "-(-4*-10)"
	actual, _ = calculate.Calculate(input)
	assert.Equal(t, "-40", actual)
}
