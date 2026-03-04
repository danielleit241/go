package repository

type UserRepository interface {
	FindAll()
	FindAllWithPagination(query string, page, limit int)
	FindById()
	Create()
	Update()
	Delete()
	IsEmailExists(email string)
}
