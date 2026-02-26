package service

import (
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

func (us *userService) GetAllUsers() ([]models.User, error) {
	return []models.User{}, nil
}

func (us *userService) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}

func (us *userService) CreateUser(user models.User) (*models.User, error) {
	user.Email = utils.NormalizeString(user.Email)

	existingUser := us.inMemoryRepo.IsEmailExists(user.Email)
	if existingUser {
		return nil, utils.NewError("email already exist", utils.ErrCodeConflict)
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
