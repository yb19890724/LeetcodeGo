package other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//
func TestAccumulated(t *testing.T) {
	
	assert.Equal(t, WaterVolume([]int{0,1,0,2,1,0,1,3,2,1,2,1}), 6, " accumulated error")
}