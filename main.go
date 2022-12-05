package main

import (
	"fmt"
	"food_mlem/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	//mongodb.CreateAcountAutomatic(6)
	//mongodb.CreateRestaurantAutomatic(10)
	//mongodb.CreateFoodAutomatic(10)
	//mongodb.CreatedrinkAutomatic(10)
	//mongodb.CreateUserAutomatic(10)
	//filter := bson.M{
	//	"user_name": "cuongmv0",
	//	"pass_word": "Cuong!@# 0",
	//}
	//acount.FindAcount(filter)
	//fmt.Println(acount.FindAcount(filter))
	app := fiber.New()
	router.AppRegister(app)
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user", "admin")
		return c.Next()
	})

	app.Get("/admin", func(c *fiber.Ctx) error {
		if c.Locals("user") == "admin" {
			return c.Status(fiber.StatusOK).SendString("Welcome, admin!")
		}
		return c.SendStatus(fiber.StatusForbidden)

	})
	app.Get("/protected", func(c *fiber.Ctx) error {
		claimData := c.Locals("jwtClaims")
		fmt.Println("claim", claimData)
		if claimData == nil {
			return c.SendString("Jwt was bypassed")
		}
		fmt.Println("claim", claimData)
		return c.JSON(claimData)

	})
	errA := app.Listen(":8080")
	if errA != nil {
		log.Println("SERVICE START ERROR: " + errA.Error())
	} else {
		log.Println("SERVICE RUNNING ON PORT 9595")
	}
}
