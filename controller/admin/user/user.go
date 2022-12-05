package user

import (
	"fmt"
	"food_mlem/model/mongodb"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(ctx *fiber.Ctx) error {
	body := new(mongodb.User)
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.JSON(map[string]interface{}{
			"message": " missing param",
		})
	}
	user := new(mongodb.User)
	filter := bson.M{
		"user_name": body.UserName,
	}
	u, er := user.FindOneUser(filter)
	if er != nil {
		return ctx.JSON(map[string]interface{}{
			"message": "can't find user" + er.Error(),
		})
	}
	fmt.Println(u)
	return ctx.JSON(map[string]interface{}{
		"message": "success",
	})

}
