package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodeDecorator_Generate(t *testing.T) {
	cd := NewCodeDecorator(&BaseCode{}, "partB")
	code := cd.Generate()
	assert.Equal(t, code, "partApartB")
}
