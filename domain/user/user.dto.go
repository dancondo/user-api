package user

import (
	userRepository "github.com/dancondo/users-api/repository/user-repository"
)

// UserDto
// contains all the data about an user
type UserDto struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserRequestDto
// contains the necessary data to login and create an user
type UserRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginUserResponseDto
// contains the user information necessary to authenticate
type LoginUserResponseDto struct {
	ID       string `json:"id"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

func NewUserDtoFromEntity(e *userRepository.UserEntity) *UserDto {
	dto := &UserDto{
		Username: e.Username,
		Password: e.Password,
	}

	if e.ID != nil {
		dto.ID = e.ID.Hex()
	}

	return dto
}
func (u *UserDto) ToUserResponseDto(token string) *LoginUserResponseDto {
	return &LoginUserResponseDto{
		ID:       u.ID,
		Token:    token,
		Username: u.Username,
	}
}
