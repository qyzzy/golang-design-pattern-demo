package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_LoginOrRegister(t *testing.T) {
	service := UserService{}
	user, err := service.LoginOrRegister("a", "b")
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}
