package contracts

type UserRepository interface {
	FindByUsername(username string) (*User, error)
}
