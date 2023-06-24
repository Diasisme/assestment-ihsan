package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ihsansolusi.co.id/bcas/service-api-go/handler/contract"
	"ihsansolusi.co.id/bcas/service-api-go/models"
	"ihsansolusi.co.id/bcas/service-api-go/pkg/logging"
)


var	DB *gorm.DB
var	log *logging.Logger


func NasabahHandlerCreate(ctx *fiber.Ctx, userData *models.User) (err error) {
	user := new(contract.RegisterRequest)
	if err := ctx.BodyParser(user); err != nil {
		log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, "Failed to do this")
		return err
	}

	errCheckNasabah := DB.Find(&userData).Error
	if errCheckNasabah != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Data NIK dan Nomor HP sudah tersedia",
		})
	}


	errCreateNasabah := DB.Create(&userData).Error
	if errCreateNasabah != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "gagal menyimpan data",
		})
	}
	
	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
	})

}

