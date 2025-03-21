package users

type IUserService interface{}

type userService struct{}

func New() IUserService {
	return &userService{}
}
