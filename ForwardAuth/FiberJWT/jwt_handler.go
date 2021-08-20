package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func handle_when_valid_token(c *fiber.Ctx) error {
	fmt.Println("***handle_when_valid_token")
	fmt.Println("BaseURL", c.BaseURL())
	fmt.Println("c.Route.Path", c.Route().Path)
	fmt.Println("c.Request.URI", c.Request().URI())
	c.Next()
	return nil
}

func handle_when_invalid_token(c *fiber.Ctx, e error) error {
	fmt.Println("Error when parsing JWT", e)
	return e
}
