package route

import (
	"github.com/gofiber/fiber/v2"
	"ihsansolusi.co.id/bcas/service-api-go/handler"
)

func RouteInit(r *fiber.App)  {
	r.Post("/daftar", handler.NasabahHandlerCreate)
}
