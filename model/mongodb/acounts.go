package mongodb

import (
	"context"
	"food_mlem/connections/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Acount struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Role     string             `bson:"role" json:"role"`
	UserName string             `bson:"user_name" json:"user_name"`
	PassWord string             `bson:"pass_word" json:"pass_word"`
}

var AcountCollection = mongodb.Mongodb.Db.Collection("acounts")

func (a *Acount) FindOneAcount(filter bson.M) (Acount, error) {
	var acount Acount
	err := userCollection.FindOne(context.TODO(), filter).Decode(&acount)
	return acount, err
}
