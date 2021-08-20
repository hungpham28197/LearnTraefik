package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte("TuyệtMật")

func main() {
	app := fiber.New()

	// Login route
	app.Post("/login", login)

	// Unauthenticated route
	app.Get("/", accessible)

	// JWT Middleware sẽ chặn trước tất các request đê
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:     SecretKey,
		SuccessHandler: handle_when_valid_token,
		ErrorHandler:   handle_when_invalid_token,
	}))

	// Restricted Routes
	app.Get("/restricted", restricted)

	//Authenticate JWT sử dụng cho Traefik ForwardAuth
	app.Get("/auth", authenticate)

	app.Listen(":3000")
}

func login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["role"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString(SecretKey)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}

func authenticate(c *fiber.Ctx) error {
	fmt.Println("***authenticate")
	fmt.Println(c)
	fmt.Println("BaseURL", c.BaseURL())
	fmt.Println("c.Request.URI", c.Request().URI())
	return c.Status(200).SendString("Authenticated!")
}
