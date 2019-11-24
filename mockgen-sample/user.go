package mockgen_sample

func NewUser(username, password, role string) *user {
	return &user{
		username: username,
		password: password,
		role:     role,
	}
}

func (u user) Name() string {
	return u.username
}

func (u user) Role() string {
	return u.role
}

func (u *user) ChangePassword(pw string) error {
	return nil
}
