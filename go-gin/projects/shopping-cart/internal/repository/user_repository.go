package repository

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (repo *userRepository) FindAll() {

}

func (repo *userRepository) FindAllWithPagination(query string, page, limit int) {
}

func (repo *userRepository) FindById() {

}

func (repo *userRepository) Create() {
}

func (repo *userRepository) Update() {

}

func (repo *userRepository) Delete() {
}

func (repo *userRepository) IsEmailExists(email string) {

}
