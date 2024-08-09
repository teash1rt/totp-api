package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pquerna/otp/totp"
)

func main() {
	key := os.Getenv("KEY_2FA")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		code, err := totp.GenerateCode(key, time.Now())
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"err_msg": "generate err"})
		}
		return c.Status(200).JSON(fiber.Map{"key": code})
	})

	log.Fatal(app.Listen(":9000"))
}
