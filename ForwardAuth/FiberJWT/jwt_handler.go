package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func handle_when_valid_token(c *fiber.Ctx) error {
	fmt.Println(c.Get("Authorization")) //Trích xuất ra đoạn header
	c.Next()
	return nil
}

func handle_when_invalid_token(c *fiber.Ctx, e error) error {
	fmt.Println("Error when parsing JWT", e)
	return e
}
