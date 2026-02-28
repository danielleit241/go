package service

import (
	"fmt"

	"github.com/danielleit241/internal/models"
	"github.com/danielleit241/internal/repository"
	"github.com/danielleit241/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	inMemoryRepo repository.UserRepository
}

func NewUserService(inMemoryRepo repository.UserRepository) UserService {
	return &userService{
		inMemoryRepo: inMemoryRepo,
	}
}

func (us *userService) GetAllUsersWithPagination(query string, page, limit int) (data []models.User, total int, err error) {
	users, total, err := us.inMemoryRepo.FindAllWithPagination(query, page, limit)
	if err != nil {
		return nil, 0, utils.WrapError("failed to retrieve users with pagination", utils.ErrCodeInternalServerError, err)
	}
	return users, total, nil
}

func (us *userService) GetAllUsers() ([]models.User, error) {
	users, err := us.inMemoryRepo.FindAll()
	if err != nil {
		return nil, utils.WrapError("failed to retrieve users", utils.ErrCodeInternalServerError, err)
	}

	return users, nil
}

func (us *userService) GetUserByID(id uuid.UUID) (*models.User, error) {
	user, err := us.inMemoryRepo.FindById(id)
	if err != nil {
		return nil, utils.WrapError(fmt.Sprintf("failed to retrieve user with id [%s]", id), utils.ErrCodeInternalServerError, err)
	}

	return user, nil
}

func (us *userService) CreateUser(user models.User) (*models.User, error) {
	user.Email = utils.NormalizeString(user.Email)

	existingUser := us.inMemoryRepo.IsEmailExists(user.Email)
	if existingUser {
		message := fmt.Sprintf("email [%s] already exists", user.Email)
		return nil, utils.NewError(message, utils.ErrCodeConflict)
	}

	user.ID = uuid.New()

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, utils.WrapError("failed to hash password", utils.ErrCodeInternalServerError, err)
	}
	user.Password = string(bcryptPassword)

	createdUser, err := us.inMemoryRepo.Create(user)
	if err != nil {
		return nil, utils.WrapError("failed to create user", utils.ErrCodeInternalServerError, err)
	}
	return createdUser, nil
}

func (us *userService) UpdateUser(id int, user models.User) (*models.User, error) {
	return nil, nil
}

func (us *userService) DeleteUser(id int) error {
	return nil
}
