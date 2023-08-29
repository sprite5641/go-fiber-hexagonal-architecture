package domain

type User struct {
	ID   string
	Name string
	Age  int
}

type UserRepository interface {
	Save(User) error
	FindByID(string) (*User, error)
}
