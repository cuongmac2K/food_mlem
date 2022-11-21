package mongodb

import (
	"context"
	"food_mlem/connections/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Drink struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	IdDrink   string             `bson:"id_drink" json:"id_drink"`
	Image     string             `bson:"image" json:"image"`
	NameDrink string             `bson:"name_drink" json:"name_drink"`
	Price     string             `bson:"price" json:"price"`
}

var drinkCollection = mongodb.Mongodb.Db.Collection("drinks")

func (d *Drink) FindOneDrink(filter bson.M) (Drink, error) {
	var drink Drink
	err := userCollection.FindOne(context.TODO(), filter).Decode(&drink)
	return drink, err
}
