package mockgen_sample

type userManagement struct {
	storage map[string]User
}

func NewUserManagement() *userManagement {
	um := &userManagement{
		storage: make(map[string]User, 0),
	}
	return um
}

func (um *userManagement) ListUsers() ([]User, error) {
	if len(um.storage) == 0 {
		return nil, EmptyStorageError{}
	}

	ul := make([]User, 0)
	for _, cur := range um.storage {
		ul = append(ul, cur)
	}

	return ul, nil
}

func (um *userManagement) CreateUser(u User) error {
	if _, ok := um.storage[u.Username()]; ok {
		return UserAlreadyExistsError{}
	}

	um.storage[u.Username()] = u
	return nil
}

func (um *userManagement) ReadUser(username string) (User, error) {
	if _, ok := um.storage[username]; !ok {
		return nil, UserDoesNotExistError{}
	}

	u := um.storage[username]
	return u, nil
}

func (um *userManagement) UpdateUser(u User) error {
	if _, ok := um.storage[u.Username()]; !ok {
		return UserDoesNotExistError{}
	}

	um.storage[u.Username()] = u
	return nil
}

func (um *userManagement) DeleteUser(username string) error {
	if _, ok := um.storage[username]; !ok {
		return UserDoesNotExistError{}
	}

	delete(um.storage, username)
	return nil
}
