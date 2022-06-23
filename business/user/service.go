package user

import (
	"errors"

	"github.com/mashbens/cps/business/user/entity"
)

type UserRepository interface {
	InsertUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	ResetPassword(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByUserID(userID string) (entity.User, error)
	UpdateUserExpiry(userID string, expiry string, memberType string) error
}

type UserService interface {
	CreateUser(user entity.User) (*entity.User, error)
	UpdateUser(user entity.User) (*entity.User, error)
	ResetPassword(user entity.User) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	FindUserByID(userID string) (*entity.User, error)
	UpdateUserExpiry(userID string, expiry string, memberType string) error
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(
	userRepo UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (c *userService) CreateUser(user entity.User) (*entity.User, error) {

	u, err := c.userRepo.FindByEmail(user.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	u, err = c.userRepo.InsertUser(user)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (c *userService) FindUserByEmail(email string) (*entity.User, error) {
	user, err := c.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *userService) ResetPassword(user entity.User) (*entity.User, error) {

	u, err := c.userRepo.ResetPassword(user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &u, nil
}

func (c *userService) FindUserByID(userID string) (*entity.User, error) {
	user, err := c.userRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *userService) UpdateUser(user entity.User) (*entity.User, error) {

	u, err := c.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (c *userService) UpdateUserExpiry(userID string, expiry string, memberType string) error {

	user := c.userRepo.UpdateUserExpiry(userID, expiry, memberType)
	return user
}
