package main

import (
	"fmt"
	"food_mlem/connections/mongodb"
)

func main() {
	mongodb.CreateAcountAutomatic(3)
	//mongodb.CreateRestaurantAutomatic(10)
	//mongodb.CreateFoodAutomatic(10)
	//mongodb.CreatedrinkAutomatic(10)
	//mongodb.CreateUserAutomatic(10)
	fmt.Println(mongodb.Mongodb.Db.Name())

}
