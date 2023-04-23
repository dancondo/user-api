package userRepository

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       *primitive.ObjectID `bson:"_id,omitempty"`
	Username string              `bson:"username"`
	Password string              `bson:"password"`
}
