package users

type IUserHandler interface{}

type userHandler struct{}

func New() IUserHandler {
	return &userHandler{}
}
