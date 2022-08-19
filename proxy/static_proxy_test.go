package proxy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserProxy_Login(t *testing.T) {
	proxy := NewUserProxy(&User{})

	err := proxy.Login("10086", "password")

	require.Nil(t, err)
}
