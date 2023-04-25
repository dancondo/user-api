package userRepository

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var (
	createMockEntity = UserEntity{
		Password: "xpto",
		Username: "john.coltrane",
	}
	mockEntity = UserEntity{
		ID:       &primitive.ObjectID{},
		Password: createMockEntity.Password,
		Username: createMockEntity.Username,
	}
	mockDBEntity = bson.D{
		primitive.E{Key: "_id", Value: mockEntity.ID},
		primitive.E{Key: "username", Value: mockEntity.Username},
		primitive.E{Key: "password", Value: mockEntity.Password},
	}
)

func Test_usersRepository_FindOneByUsername(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	type args struct {
		username string
	}
	tests := []struct {
		name       string
		r          *usersRepository
		args       args
		want       *UserEntity
		wantErr    bool
		setupMTest func(mt *mtest.T) *mongo.Collection
	}{
		{
			name: "Should return an user if found",
			r:    &usersRepository{},
			args: args{
				username: mockEntity.Username,
			},
			want:    &mockEntity,
			wantErr: false,
			setupMTest: func(mt *mtest.T) *mongo.Collection {
				mockCollection := mt.Coll

				mt.AddMockResponses(mtest.CreateCursorResponse(1, "users-api-db.users", mtest.FirstBatch, mockDBEntity))

				return mockCollection
			},
		},
		{
			name: "Should not return an user if not found",
			r:    &usersRepository{},
			args: args{
				username: mockEntity.Username,
			},
			want:    nil,
			wantErr: false,
			setupMTest: func(mt *mtest.T) *mongo.Collection {
				mockCollection := mt.Coll

				mt.AddMockResponses(mtest.CreateCursorResponse(0, "users-api-db.users", mtest.FirstBatch))

				return mockCollection
			},
		},
		{
			name: "Should return an error if error",
			r:    &usersRepository{},
			args: args{
				username: mockEntity.Username,
			},
			want:    nil,
			wantErr: true,
			setupMTest: func(mt *mtest.T) *mongo.Collection {
				mockCollection := mt.Coll

				mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
					Code:    1,
					Message: "Error",
				}))

				return mockCollection
			},
		},
	}
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			tt.r.collection = tt.setupMTest(mt)
			got, err := tt.r.FindOneByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersRepository.FindOneByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersRepository.FindOneByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersRepository_Create(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	type args struct {
		user *UserEntity
	}
	tests := []struct {
		name       string
		r          *usersRepository
		args       args
		want       *UserEntity
		wantErr    bool
		setupMTest func(mt *mtest.T) *mongo.Collection
	}{
		{
			name: "Should create an user on success",
			r:    &usersRepository{},
			args: args{
				user: &createMockEntity,
			},
			want:    &createMockEntity,
			wantErr: false,
			setupMTest: func(mt *mtest.T) *mongo.Collection {
				mockCollection := mt.Coll

				mt.AddMockResponses(mtest.CreateSuccessResponse())

				return mockCollection
			},
		},
		{
			name: "Should return an error if unable to create",
			r:    &usersRepository{},
			args: args{
				user: &mockEntity,
			},
			want:    nil,
			wantErr: true,
			setupMTest: func(mt *mtest.T) *mongo.Collection {
				mockCollection := mt.Coll

				mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
					Code:    1,
					Message: "Error",
				}))

				return mockCollection
			},
		},
	}
	for _, tt := range tests {
		mt.Run(tt.name, func(mt *mtest.T) {
			tt.r.collection = tt.setupMTest(mt)
			got, err := tt.r.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usersRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
