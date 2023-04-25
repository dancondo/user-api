package user

import (
	"fmt"

	"github.com/dancondo/users-api/common"
	"github.com/dancondo/users-api/pkg/cryptography"
	userRepository "github.com/dancondo/users-api/repository/user-repository"
)

type UserService interface {
	CreateUser(user *UserRequestDto) (*UserDto, error)
	GetUserByUsername(username string) (*UserDto, error)
	ValidateUserPassword(user *UserDto, password string) error
}

type userService struct {
	crypto     cryptography.Crypto
	repository userRepository.UsersRepository
}

func NewService() UserService {
	return &userService{
		crypto:     cryptography.New(),
		repository: userRepository.New(),
	}
}

func (s *userService) CreateUser(user *UserRequestDto) (*UserDto, error) {
	password, err := s.encryptPassword(user.Password)

	if err != nil {
		return nil, err
	}

	entity, err := s.repository.Create(&userRepository.UserEntity{
		Username: user.Username,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	return NewUserDtoFromEntity(entity), nil
}

func (s *userService) GetUserByUsername(username string) (*UserDto, error) {
	entity, err := s.repository.FindOneByUsername(username)

	if err != nil {
		common.Log.Errorf("[USER SERVICE][GET USER BY USERNAME], %v", err.Error())
		return nil, err
	}

	if entity == nil {
		return nil, nil
	}

	return NewUserDtoFromEntity(entity), nil
}

func (s *userService) ValidateUserPassword(user *UserDto, password string) error {
	valid := s.crypto.ValidatePassword(user.Password, password)

	if !valid {
		return fmt.Errorf("Invalid password")
	}

	return nil
}

func (s *userService) encryptPassword(password string) (string, error) {
	return s.crypto.EncryptPassword(password)
}
