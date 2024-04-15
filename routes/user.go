package routes

import (
	"github.com/4fort/fiber_gorm/database"
	"github.com/4fort/fiber_gorm/models"
	"github.com/gofiber/fiber/v3"
)

type UserSerializer struct {
	// not the model User, this is a serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(user models.User) UserSerializer {
	return UserSerializer{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}
