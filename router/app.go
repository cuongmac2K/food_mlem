package router

import (
	"food_mlem/controller/admin/acount"
	"food_mlem/controller/admin/user"
	"food_mlem/jwtchecker"
	"github.com/gofiber/fiber/v2"
)

func AppRegister(app *fiber.App) {
	Acount(app)
	User(app)
}
func Acount(app *fiber.App) {
	ac := app.Group("/acount")
	ac.Post("/login", acount.Login)
	ac.Post("/register", acount.Register)
}
func User(app *fiber.App) {
	ac := app.Group("/user")
	ac.Use(jwtchecker.New(jwtchecker.Config{}))
	ac.Post("/get_user", user.GetUser)
}
