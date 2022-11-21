package mongodb

import (
	"context"
	"food_mlem/connections/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	IdRestaurant   string             `bson:"id_restaurant" json:"id_restaurant"`
	Image          string             `bson:"image" json:"image"`
	NameRestaurant string             `bson:"name_restaurant" json:"name_restaurant"`
	Address        string             `bson:"address" json:"address"`
}

var RestaurantCollection = mongodb.Mongodb.Db.Collection("restaurants")

func (r *Restaurant) FindOneRestaurant(filter bson.M) (Restaurant, error) {
	var restaurant Restaurant
	err := userCollection.FindOne(context.TODO(), filter).Decode(&restaurant)
	return restaurant, err
}
