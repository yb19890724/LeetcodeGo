package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinationsBT(t *testing.T) {

	valid := LetterCombinationsBT("23")

	assert.Equal(t, valid, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, " Output Error ")
}
