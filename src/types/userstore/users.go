package userstore

import "user-management/src/types"

type userManagement struct {
	storage map[string]types.User
}

func NewUserManagement() (*userManagement, error) {
	um := &userManagement{
		storage: make(map[string]types.User, 0),
	}

	err := um.initializeUserManagement()
	if err != nil {
		return nil, err
	}

	return um, nil
}

func (um *userManagement) initializeUserManagement() error {
	// this is "super safe" too
	u, err := NewUser("admin", "secretPass", "admin")
	if err != nil {
		return err
	}

	um.storage["admin"] = u
	return nil
}

func (um userManagement) ListUsers() ([]types.User, error) {
	if len(um.storage) == 0 {
		return nil, EmptyStorageError{}
	}

	ul := make([]types.User, 0)
	for _, cur := range um.storage {
		ul = append(ul, cur)
	}

	return ul, nil
}

func (um *userManagement) CreateUser(u types.User) error {
	if _, ok := um.storage[u.Username()]; ok {
		return UserAlreadyExistsError{}
	}

	um.storage[u.Username()] = u
	return nil
}

func (um userManagement) ReadUser(username string) (types.User, error) {
	if _, ok := um.storage[username]; !ok {
		return nil, UserDoesNotExistError{}
	}

	u := um.storage[username]
	return u, nil
}

func (um *userManagement) UpdateUser(u types.User) error {
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
