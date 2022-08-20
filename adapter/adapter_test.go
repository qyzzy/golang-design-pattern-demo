package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSDAdapterTF_ReadSD(t *testing.T) {
	tf := &TFCardImpl{}
	sda := &SDAdapterTF{
		tfCard: tf,
	}
	b, err := sda.ReadSD()
	assert.Equal(t, b, []byte("testtf"))
	assert.Nil(t, err)
}
