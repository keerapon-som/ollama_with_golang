package usersvc

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) GetAllUsers() ([]string, error) {
	return []string{"Alice", "Bob"}, nil
}
