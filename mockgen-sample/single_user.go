package mockgen_sample

type user struct {
	username string
	password string
	role     string
}

func NewUser(username, password, role string) *user {
	return &user{
		username: username,
		password: password,
		role:     role,
	}
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
