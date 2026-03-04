package v1service

type UserService interface {
	GetAllUsers()
	GetAllUsersWithPagination(query string, page, limit int)
	GetUserByID()
	CreateUser()
	UpdateUser()
	DeleteUser()
}
