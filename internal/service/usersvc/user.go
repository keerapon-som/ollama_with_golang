package usersvc

type UserService interface {
	GetAllUsers() ([]string, error)
}
