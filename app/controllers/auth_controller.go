package controllers

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/prabhatpankaj/go-fiber-rest-api/app/models"
	"github.com/prabhatpankaj/go-fiber-rest-api/pkg/utils"
	"github.com/prabhatpankaj/go-fiber-rest-api/platform/database"
	"golang.org/x/crypto/bcrypt"
)

// func SignUpUser(c *fiber.Ctx) error {
// 	var payload *models.SignUpInput

// 	if err := c.BodyParser(&payload); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
// 	}

// 	errors := models.ValidateStruct(payload)
// 	if errors != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})

// 	}

// 	if payload.Password != payload.PasswordConfirm {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})

// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
// 	}

// 	newUser := models.User{
// 		Name:     payload.Name,
// 		Email:    strings.ToLower(payload.Email),
// 		Password: string(hashedPassword),
// 		Photo:    &payload.Photo,
// 	}

// 	result := initializers.DB.Create(&newUser)

// 	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
// 		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists"})
// 	} else if result.Error != nil {
// 		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": models.FilterUserRecord(&newUser)}})
// }

func SignUpUser(c *fiber.Ctx) error {
	var payload *models.SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail a", "message": err.Error()})
	}

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})

	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status code and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if username or Email exists.
	foundedUsername, err := db.GetUserByUsername(payload.Username)
	log.Println(foundedUsername)
	if err == nil {
		// Return status 404 and book not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with this username is already exists",
		})
	}

	// Checking, if username or Email exists.
	foundedEmail, err := db.GetProflebyEmail(payload.Email)
	log.Println(foundedEmail)
	if err == nil {
		// Return status 404 and book not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with this email is already exists",
		})
	}

	// Create new Users struct
	user := &models.Users{}

	// Set initialized default data for book:
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.ActiveStatus = 1 // 0 == draft, 1 == active
	user.Username = strings.ToLower(payload.Username)
	user.Password = string(hashedPassword)

	// Create a new validator for a Users model.
	validate := utils.NewValidator()

	// Validate users fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create User by given ID.
	if err := db.CreateUser(user); err != nil {
		// Return status code and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// newProfile := models.Profile{
	// 	UserId:         newUser.ID,
	// 	Email:          payload.Email,
	// 	Role:           "Customer",
	// 	FullName:       payload.FullName,
	// 	VerifiedStatus: 1,
	// }

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}
