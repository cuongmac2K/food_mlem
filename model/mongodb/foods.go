package mongodb

import (
	"context"
	"food_mlem/connections/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	IdFood   string             `bson:"id_food" json:"id_food"`
	Image    string             `bson:"image" json:"image"`
	NameFood string             `bson:"name_food" json:"name_food"`
	Price    string             `bson:"price" json:"price"`
}

var FoodCollection = mongodb.Mongodb.Db.Collection("foods")

func (f *Food) FindOneFood(filter bson.M) (Food, error) {
	var food Food
	err := userCollection.FindOne(context.TODO(), filter).Decode(&food)
	return food, err
}
