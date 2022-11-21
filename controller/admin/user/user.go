package user

//func CreateUserAutomatic(n int) {
//	var userCollection = Mongodb.Db.Collection("user")
//	for i := 0; i < n; i++ {
//		id_user := 1 + i
//		user_name := "john" + string(i)
//		address := "HN" + string(i)
//		salt := ""
//		avatar := ""
//		create_at := time.Now()
//		update_at := ""
//		date_of_birth := ""
//		doc := bson.D{
//			{"id_user", id_user},
//			{"user_name", user_name},
//			{"address", address},
//			{"salt", salt},
//			{"avatar", avatar},
//			{"create_at", create_at},
//			{"update_at", update_at},
//			{"date_of_birth", date_of_birth},
//		}
//		result, err := userCollection.InsertOne(context.TODO(), doc)
//		if err != nil {
//			fmt.Println("err la insert ", err)
//		}
//		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
//	}
//
//}
