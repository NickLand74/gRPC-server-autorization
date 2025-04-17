package storage

type Storage interface {
	CreateUser(username, password string) error
	GetUser(username string) (*User, error)
}

type User struct {
	Username string
	Password string
}
