package userRepository

import (
	"context"
	"time"

	"github.com/dancondo/users-api/common"
	"github.com/dancondo/users-api/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UsersRepository manages in the database interaction in the currencies collection
type UsersRepository interface {
	CreateIndex(name string) error
	FindOneByUsername(username string) (*UserEntity, error)
	Create(user *UserEntity) (*UserEntity, error)
}

type usersRepository struct {
	collection *mongo.Collection
}

func New() UsersRepository {
	return &usersRepository{
		collection: db.GetCollection("users"),
	}
}

// FindOneByUsername finds a user in the database by username
func (r *usersRepository) FindOneByUsername(username string) (*UserEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result := r.collection.FindOne(ctx, bson.M{
		"username": username,
	})

	var user *UserEntity

	err := result.Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		common.Log.Errorf("[USERS REPOSITORY] %v", err.Error())
		return nil, err
	}

	return user, nil
}

// Create creates an user in the database
func (r *usersRepository) Create(user *UserEntity) (*UserEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)

	if err != nil {
		common.Log.Errorf("[USERS REPOSITORY] %v", err.Error())
		return nil, err
	}

	return &UserEntity{
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (r *usersRepository) CreateIndex(name string) error {
	_, err := r.collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: name, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	return err
}
