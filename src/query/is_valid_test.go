package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {

	valid := IsValid("()")

	assert.Equal(t, valid, true, " Output Error ")

	valid2 := IsValid("()[]{}")

	assert.Equal(t, valid2, true, " Output Error. ")

	valid3 := IsValid("(]")

	assert.Equal(t, valid3, false, " Output Error. ")

	valid4 := IsValid("([)]")

	assert.Equal(t, valid4, false, " Output Error. ")

	valid5 := IsValid("{[]}")

	assert.Equal(t, valid5, true, " Output Error. ")
}
