package facade

type UserInterface interface {
	Login(mobileNumber, password string) (*User, error)
	Register(mobileNumber, password string) (*User, error)
}

type UserFacadeInterface interface {
	LoginOrRegister(mobileNumber, password string) (*User, error)
}

type User struct {
	Name string
}

type UserService struct {
}

func (u *UserService) Login(mobileNumber, password string) (*User, error) {
	return &User{
		Name: "test login",
	}, nil
}

func (u *UserService) Register(mobileNumber, password string) (*User, error) {
	return &User{
		Name: "test register",
	}, nil
}

func (u *UserService) LoginOrRegister(mobileNumber, password string) (*User, error) {
	user, err := u.Login(mobileNumber, password)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, err
	}
	return u.Register(mobileNumber, password)
}
