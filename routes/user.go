package routes

import (
	"errors"
	"time"

	"github.com/4fort/fiber_gorm/database"
	"github.com/4fort/fiber_gorm/models"
	"github.com/gofiber/fiber/v3"
)

type UserSerializer struct {
	// not the model User, this is a serializer
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateResponseUser(user models.User) UserSerializer {
	return UserSerializer{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt}
}

func CreateUser(c fiber.Ctx) error {
	var user models.User

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c fiber.Ctx) error {
	var users []models.User

	database.Database.Db.Find(&users)
	responseUsers := []UserSerializer{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(user *models.User, c fiber.Ctx) error {
	id := c.Params("id")

	database.Database.Db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return errors.New(fiber.ErrNotFound.Error())
	}
	return nil
}

func GetUser(c fiber.Ctx) error {
	var user models.User

	if err := findUser(&user, c); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c fiber.Ctx) error {
	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var user models.User
	var updatedData UpdateUser

	if err := findUser(&user, c); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := c.Bind().Body(&updatedData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updatedData.FirstName
	user.LastName = updatedData.LastName

	// database.Database.Db.Save(&user)
	if err := database.Database.Db.Model(&user).Updates(&user).Error; err != nil {
		return c.Status(500).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c fiber.Ctx) error {
	var user models.User

	if err := findUser(&user, c); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(500).JSON(err.Error)
	}

	return c.Status(204).JSON("Successfully deleted user")
}
