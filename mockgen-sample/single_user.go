package mockgen_sample

type user struct {
	username string
	password string
	role     string
}

func NewUser(username, password, role string) (*user, error) {
	u := user{
		username: username,
		password: password,
		role:     role,
	}

	if err := u.validateUser(); err != nil {
		return nil, err
	}

	return &u, nil
}

func NewUserFromDTO(dto UserDTO) (*user, error) {
	return NewUser(dto.Username, dto.Password, dto.Role)
}

func (u user) Username() string {
	return u.username
}

func (u user) Role() string {
	return u.role
}

func (u *user) ChangePassword(pw string) error {
	if len(pw) < 8 {
		return &InvalidPasswordError{}
	}
	u.password = pw
	return nil
}

func (u user) validateUser() error {
	if len(u.username) == 0 {
		return InvalidUsernameError{}
	}

	if len(u.password) < 8 {
		return InvalidPasswordError{}
	}

	if u.role != "admin" && u.role != "member" {
		return InvalidRoleError{}
	}

	return nil
}
