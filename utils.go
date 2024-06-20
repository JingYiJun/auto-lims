package main

import (
	"github.com/gofiber/fiber/v2"
)

func checkApikey(apikey string) (username string, err error) {
	for k, v := range keys {
		if v == apikey {
			return k, nil
		}
	}
	return "", fiber.NewError(401, "no auth")
}
