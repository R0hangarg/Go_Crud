package controllers

import (
	"Crud_fiber_Go/models"
	"Crud_fiber_Go/services"
	"Crud_fiber_Go/utils"
	"Crud_fiber_Go/views"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUsers(c *fiber.Ctx) error {
	// Example: Fetch users from database
	users, err := services.GetUsers()
	if err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	fmt.Println(user)
	if err := c.BodyParser(user); err != nil {
		return views.JSON(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Save user to database
	// DB logic here...
	err := services.CreateUser(user)

	if err == nil {
		return c.Status(201).JSON(user)
	} else {
		return c.Status(500).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id") // Get the user ID from the URL parameter

	// Delete the user from the database
	err := services.DeleteUser(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	// if err := models.DB.First(&user, id).Error; err != nil {
	// 	return c.Status(404).JSON(fiber.Map{
	// 		"error": "User not found",
	// 	})
	// }

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	idInt, err := strconv.Atoi(id)
	// Save the updated user to the database
	err = services.UpdateUser(uint(idInt), &user)

	if err == nil {
		return c.Status(200).JSON(user)
	} else {
		return c.Status(500).JSON(map[string]interface{}{
			"Error": err.Error(),
		})
	}
}

func LoginUser(c *fiber.Ctx) error {
	// Parse the request body into a Login model
	user := new(models.Login)

	if err := c.BodyParser(user); err != nil {
		return views.JSON(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate input
	if user.Email == "" || user.Password == "" {
		return views.JSON(c, fiber.StatusBadRequest, "Email and Password are required")
	}

	// Check if the user exists in the database
	var dbUser models.User
	if err := models.DB.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.JSON(c, fiber.StatusUnauthorized, "Invalid email")
		}
		return views.JSON(c, fiber.StatusInternalServerError, "Database error")
	}

	matchPassword := utils.CheckPasswordHash(user.Password, dbUser.Password)
	// Compare the password (in a real-world app, you would hash and compare passwords)
	if !matchPassword {
		return views.JSON(c, fiber.StatusUnauthorized, "Invalid password")
	}

	token, _ := utils.GenerateJWT(user.Email)
	// Respond with success (e.g., JWT token in a real-world app)
	return views.JSON(c, fiber.StatusOK, fiber.Map{
		"message": "Login successful",
		"user":    dbUser,
		"token":   token,
	})
}
