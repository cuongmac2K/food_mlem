package acount

import (
	"food_mlem/common"
	"food_mlem/jwtchecker"
	"food_mlem/model/mongodb"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

//func FindAcount(ctx *fiber.Ctx) error {
//	body := new(Acount)
//	if err := ctx.BodyParser(&body); err != nil {
//		return ctx.JSON(map[string]interface{}{
//			"data": " missing param",
//		})
//	}
//	return ctx.JSON(map[string]interface{}{
//		"data": " chay r nha",
//	})
//
//}
func Login(ctx *fiber.Ctx) error {
	body := new(LoginBody)
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.JSON(map[string]interface{}{
			"data": " missing param",
		})
	}
	ac := new(mongodb.Acount)
	filter := bson.M{
		"user_name": body.UserName,
		"pass_word": body.PassWord,
	}
	_, err := ac.FindOneAcount(filter)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"message": "ko thay user",
			"error":   401,
		})
	}

	token, err := jwtchecker.Encode(&jwt.MapClaims{
		"u":    "test@gmail.com",
		"role": "user",
	}, 1000)

	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"message": "err with jwt" + err.Error(),
		})
	}
	return ctx.JSON(map[string]interface{}{
		"message": "user exist",
		"token":   token,
	})
}
func Register(ctx *fiber.Ctx) error {
	body := new(RegisterBody)
	if err := ctx.BodyParser(&body); err != nil || body.PassWord == "" || body.UserName == "" {
		return ctx.JSON(map[string]interface{}{
			"data": " missing param",
		})
	}
	ac := new(mongodb.Acount)
	salts := common.GenSalt(50)
	pass := common.Hash(body.PassWord + salts)
	filter := bson.M{
		"user_name": body.UserName,
		"pass_word": pass,
		"role":      body.Role,
	}
	err := ac.InsertAcount(filter)
	if err != nil {
		return ctx.JSON(map[string]interface{}{
			"message": "ko inster acount dc" + err.Error(),
			"error":   401,
		})
	}
	return ctx.JSON(map[string]interface{}{
		"message": "create acount success",
	})
}
