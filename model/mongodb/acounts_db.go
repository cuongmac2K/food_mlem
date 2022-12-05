package mongodb

import (
	"context"
	"fmt"
	"food_mlem/connections/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Acount struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Role     string             `bson:"role" json:"role"`
	UserName string             `bson:"user_name" json:"user_name"`
	PassWord string             `bson:"pass_word" json:"pass_word"`
	Status   int                `bson:"status" json:"status"`
}

var AcountCollection = mongodb.Mongodb.Db.Collection("acounts")

func (a *Acount) FindOneAcount(filter bson.M) (Acount, error) {
	var acount Acount
	err := AcountCollection.FindOne(context.TODO(), filter).Decode(&acount)
	return acount, err
}
func (a *Acount) InsertAcount(filter bson.M) error {
	result, err := AcountCollection.InsertOne(context.TODO(), filter)
	fmt.Println("insert result ", result)
	return err
}
