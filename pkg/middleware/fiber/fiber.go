package fiberopenapiui

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/openapi-ui/go-openapi-ui/pkg/doc"
)

func New(doc doc.Doc) fiber.Handler {
	return adaptor.HTTPHandlerFunc(doc.Handler())
}
