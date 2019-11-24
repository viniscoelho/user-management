//go:generate mockgen -destination=mocks/mocks.go -package=mocks tdc-presentation/mockgen-sample Users
package mockgen_sample

type user struct {
	username string
	password string
	role     string
}

type UserDTO struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

type User interface {
	Username() string
	Role() string
	ChangePassword(pw string) error
}

type Users interface {
	ListUsers() ([]Users, error)
	CreateUser(u User) error
	ReadUser() (User, error)
	UpdateUser(u User) error
	DeleteUser(username string) error
}
