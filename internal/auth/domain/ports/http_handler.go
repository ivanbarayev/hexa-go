package ports

import (
	"github.com/gofiber/fiber/v2"
)

// IHandlers Auth Domain HTTP handler interface
type IHandlers interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
