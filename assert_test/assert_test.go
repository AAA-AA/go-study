package asserttest_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOut(t *testing.T) {
	ret := 1
	assert.Equal(t, ret, 1)
}
