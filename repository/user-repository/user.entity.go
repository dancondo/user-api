package userRepository

type UserEntity struct {
	ID       *interface{} `bson:"_id,omitempty"`
	Username string       `bson:"username"`
	Password string       `bson:"password"`
}
