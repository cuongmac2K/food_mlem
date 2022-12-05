package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"time"
)

func CreateUserAutomatic(n int) {
	var userCollection = Mongodb.Db.Collection("user")
	for i := 0; i < n; i++ {
		id_user := 1 + i
		user_name := "john" + string(i)
		address := "HN" + string(i)
		salt := ""
		avatar := ""
		create_at := time.Now()
		update_at := ""
		date_of_birth := ""
		doc := bson.D{
			{"id_user", id_user},
			{"user_name", user_name},
			{"address", address},
			{"salt", salt},
			{"avatar", avatar},
			{"create_at", create_at},
			{"update_at", update_at},
			{"date_of_birth", date_of_birth},
		}
		result, err := userCollection.InsertOne(context.TODO(), doc)
		if err != nil {
			fmt.Println("err la insert ", err)
		}
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}

}
func CreatedrinkAutomatic(n int) {
	var drinkCollection = Mongodb.Db.Collection("drinks")
	for i := 0; i < n; i++ {
		id_drink := 1 + i
		name_drink := "coca cola" + string(i)
		image := ""
		price := "1000" + string(i)
		doc := bson.D{
			{"id_drink", id_drink},
			{"name_drink", name_drink},
			{"price", price},
			{"image", image},
		}
		result, err := drinkCollection.InsertOne(context.TODO(), doc)
		if err != nil {
			fmt.Println("err la insert ", err)
		}
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}
}
func CreateFoodAutomatic(n int) {
	var drinkCollection = Mongodb.Db.Collection("foods")
	for i := 0; i < n; i++ {
		id_food := 1 + i
		name_food := "ga ran " + strconv.Itoa(i)
		image := ""
		price := "1000" + strconv.Itoa(i)
		doc := bson.D{
			{"id_food", id_food},
			{"name_food", name_food},
			{"price", price},
			{"image", image},
		}
		result, err := drinkCollection.InsertOne(context.TODO(), doc)
		if err != nil {
			fmt.Println("err la insert ", err)
		}
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}
}
func CreateRestaurantAutomatic(n int) {
	var drinkCollection = Mongodb.Db.Collection("restaurants")
	for i := 0; i < n; i++ {
		id_restaurant := 1 + i
		name_restaurant := "ga ran " + strconv.Itoa(i)
		image := ""
		address := "10" + strconv.Itoa(i) + "Le loi Stress Ha noi Viet Name"
		doc := bson.D{
			{"id_restaurant", id_restaurant},
			{"name_restaurant", name_restaurant},
			{"image", image},
			{"address", address},
		}
		result, err := drinkCollection.InsertOne(context.TODO(), doc)
		if err != nil {
			fmt.Println("err la insert ", err)
		}
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}
}
func CreateAcountAutomatic(n int) {
	var acountCollection = Mongodb.Db.Collection("acounts")
	for i := 0; i < n; i++ {
		//role := "shiper"
		//role := "food maker"
		role := "user"
		user_name := "hama" + strconv.Itoa(i)
		pass_word := "Hama!@# " + strconv.Itoa(i)
		doc := bson.D{
			{"role", role},
			{"status", 1},
			{"user_name", user_name},
			{"pass_word", pass_word},
		}
		result, err := acountCollection.InsertOne(context.TODO(), doc)
		if err != nil {
			fmt.Println("err la insert ", err)
		}
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}
}
