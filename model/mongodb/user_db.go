package mongodb

import (
	"context"
	"food_mlem/connections/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	IdUser      int                `bson:"id_user" json:"id_user"`
	UserName    string             `bson:"user_name" json:"user_name"`
	Address     string             `bson:"address" json:"address"`
	Salt        string             `bson:"salt" json:"salt"`
	Avatar      string             `bson:"avatar" json:"avatar"`
	CreateAt    time.Time          `bson:"create_at" json:"create_at"`
	UpdateAt    time.Time          `bson:"-,omitempty" json:"_id,omitempty"`
	DateOfBirth string             `bson:"date_of_birth" json:"date_of_birth"`
}

var userCollection = mongodb.Mongodb.Db.Collection("user")

func (u *User) FindOneUser(filter bson.M) (User, error) {
	var user User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}
