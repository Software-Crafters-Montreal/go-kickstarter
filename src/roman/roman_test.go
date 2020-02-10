package roman

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertOne(t *testing.T) {
    res := Translate(1)
	assert.Equal(t, res, "I", "The two words should be the same.")
}
