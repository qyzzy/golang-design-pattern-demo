package proxy

type UserInterface interface {
	Login(mobileNumber, password string) error
}

type User struct {
	mobileNumber, password string
}

func (u *User) Login(mobileNumber, password string) error {
	// Ignore details
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

func (p *UserProxy) Login(mobileNumber, password string) error {
	// Business logic
	if err := p.user.Login(mobileNumber, password); err != nil {
		return err
	}
	// Monitoring logic
	return nil
}
