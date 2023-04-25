package user

import (
	"fmt"
	"reflect"
	"testing"

	userRepository "github.com/dancondo/users-api/repository/user-repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	mockEntity = userRepository.UserEntity{
		Password: "xpto",
		Username: "john.coltrane",
		ID:       &primitive.ObjectID{},
	}
	errMock        = fmt.Errorf("Error")
	mockUserReqDto = UserRequestDto{
		Password: "xpto",
		Username: "john.coltrane",
	}
	mockUserDto = UserDto{
		ID:       mockEntity.ID.Hex(),
		Username: mockEntity.Username,
		Password: mockEntity.Password,
	}
	mockService = &userService{}
)

type mockRepo struct {
	wantErr bool
	success bool
}

func (r *mockRepo) FindOneByUsername(username string) (*userRepository.UserEntity, error) {
	if r.wantErr {
		return nil, errMock
	}

	if !r.success {
		return nil, nil
	}

	return &mockEntity, nil
}

func (r *mockRepo) Create(user *userRepository.UserEntity) (*userRepository.UserEntity, error) {
	if r.wantErr {
		return nil, errMock
	}

	return &mockEntity, nil
}

func (r *mockRepo) CreateIndex(name string) error {
	if r.wantErr {
		return errMock
	}

	return nil
}

type mockCrypto struct {
	wantErr bool
}

func (c *mockCrypto) EncryptPassword(password string) (string, error) {
	if c.wantErr {
		return "", errMock
	}

	return password, nil
}

func (c *mockCrypto) ValidatePassword(value string, comparation string) bool {
	if !c.wantErr {
		return true
	}

	return false
}

func Test_userService_CreateUser(t *testing.T) {
	type args struct {
		user *UserRequestDto
	}
	tests := []struct {
		name        string
		s           *userService
		args        args
		want        *UserDto
		wantErr     bool
		repoWantErr bool
		repoSuccess bool
	}{
		{
			name: "Should create and return an user if success",
			s:    mockService,
			args: args{
				user: &mockUserReqDto,
			},
			want:        &mockUserDto,
			wantErr:     false,
			repoWantErr: false,
			repoSuccess: true,
		},
		{
			name: "Should return an error if failed",
			s:    mockService,
			args: args{
				user: &mockUserReqDto,
			},
			want:        nil,
			wantErr:     true,
			repoWantErr: true,
			repoSuccess: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.repository = &mockRepo{wantErr: tt.repoWantErr, success: tt.repoSuccess}
			tt.s.crypto = &mockCrypto{wantErr: false}
			got, err := tt.s.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetUserByUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name        string
		s           *userService
		args        args
		want        *UserDto
		wantErr     bool
		repoWantErr bool
		repoSuccess bool
	}{
		{
			name: "Should return an user if success",
			s:    mockService,
			args: args{
				username: mockUserReqDto.Username,
			},
			want:        &mockUserDto,
			wantErr:     false,
			repoWantErr: false,
			repoSuccess: true,
		}, {
			name: "Should not return an user if not found",
			s:    mockService,
			args: args{
				username: mockUserReqDto.Username,
			},
			want:        nil,
			wantErr:     false,
			repoWantErr: false,
			repoSuccess: false,
		},
		{
			name: "Should return an error if failed",
			s:    mockService,
			args: args{
				username: mockUserReqDto.Username,
			},
			want:        nil,
			wantErr:     true,
			repoWantErr: true,
			repoSuccess: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.repository = &mockRepo{wantErr: tt.repoWantErr, success: tt.repoSuccess}
			tt.s.crypto = &mockCrypto{wantErr: false}
			got, err := tt.s.GetUserByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUserByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_ValidateUserPassword(t *testing.T) {
	type args struct {
		user     *UserDto
		password string
	}
	tests := []struct {
		name          string
		s             *userService
		args          args
		wantErr       bool
		cryptoWantErr bool
	}{
		{
			name: "Should not return an error if success",
			s:    mockService,
			args: args{
				user:     &mockUserDto,
				password: mockUserDto.Password,
			},
			wantErr:       false,
			cryptoWantErr: false,
		},
		{
			name: "Should return an error if failed",
			s:    mockService,
			args: args{
				user:     &mockUserDto,
				password: "",
			},
			wantErr:       true,
			cryptoWantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.repository = &mockRepo{wantErr: false, success: true}
			tt.s.crypto = &mockCrypto{wantErr: tt.cryptoWantErr}
			if err := tt.s.ValidateUserPassword(tt.args.user, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("userService.ValidateUserPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
