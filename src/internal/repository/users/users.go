package users

type IUserRepository interface{}

type userRepository struct{}

func New() IUserRepository {
	return &userRepository{}
}
