package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorNotification_Notify(t *testing.T) {
	emailSender := NewEmailMessageSender([]string{"hello", "qyzzy"})
	errorNotification := NewErrorNotification(emailSender)
	err := errorNotification.Notify("msg")
	assert.Nil(t, err)
}

func TestNormalNotification_Notify(t *testing.T) {
	mobileSender := NewMobileMessageSender("hello")
	normalNotification := NewNormalNotification(mobileSender)
	err := normalNotification.Notify("msg")
	assert.Nil(t, err)
}
