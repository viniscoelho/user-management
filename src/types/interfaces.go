//go:generate mockgen -destination=mocks/mocks.go -package=mocks user-management/src/types User,Users
package types

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}

type User interface {
	Username() string
	Role() string
	ChangePassword(pw string) error
}

type Users interface {
	ListUsers() ([]User, error)
	CreateUser(u User) error
	ReadUser(username string) (User, error)
	UpdateUser(u User) error
	DeleteUser(username string) error
}
