package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("error on getting hostname: %s", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hi from %s!", hostname))
	})

	log.Fatal(app.Listen(":8080"))
}
