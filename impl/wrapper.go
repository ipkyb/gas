package impl

import (
	"github.com/ipkyb/gas/api"

	"github.com/gofiber/fiber/v2"
)

func FiberHandler(handler api.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ictx := cFiber{ctx}
		handler(&ictx)
		return nil
	}
}

func FiberHandlers(handlers []api.Handler) []fiber.Handler {
	res := make([]fiber.Handler, len(handlers))
	for i, h := range handlers {
		res[i] = FiberHandler(h)
	}
	return res
}
