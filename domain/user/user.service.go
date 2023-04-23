package user

import (
	"github.com/dancondo/users-api/common"
	userRepository "github.com/dancondo/users-api/repository/user-repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *UserRequestDto) (*UserDto, error)
	GetUserByUsername(username string) (*UserDto, error)
	ValidateUserPassword(user *UserDto, password string) error
}

type userService struct {
	repository userRepository.UsersRepository
}

func NewService() UserService {
	return &userService{
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
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func (s *userService) encryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
