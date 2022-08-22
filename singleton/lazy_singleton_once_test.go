package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLazySingletonOnce_GetInstance(t *testing.T) {
	assert.Equal(t, lazySingletonOnce.GetInstance(), lazySingletonOnce.GetInstance())
}
