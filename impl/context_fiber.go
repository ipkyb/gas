package impl

import (
	"github.com/ipkyb/gas/api"

	"github.com/gofiber/fiber/v2"
)

var _ api.Context = &cFiber{}

type cFiber struct {
	ctx *fiber.Ctx
}

// Next implements Context.
func (i *cFiber) Next() {
	i.ctx.Next()
}

// SetHeader implements Context.
func (i *cFiber) SetHeader(key string, value string) api.Context {
	i.ctx.Context().Response.Header.Set(key, value)
	return i
}

// Status implements Context.
func (i *cFiber) Status(status int) api.Context {
	i.ctx.Context().Response.SetStatusCode(status)
	return i
}

// WriteInplace implements Context.
func (i *cFiber) WriteInplace(body []byte) {
	i.ctx.Context().Response.SetBodyRaw(body)
}

// WriteBytes implements Context.
func (i *cFiber) WriteBytes(body []byte) {
	i.ctx.Context().Response.SetBody(body)
}

// WriteJSON implements Context.
func (i *cFiber) WriteJSON(body interface{}) {
	i.ctx.JSON(body)
}

// WriteString implements Context.
func (i *cFiber) WriteString(body string) {
	i.ctx.Context().Response.SetBodyString(body)
}
