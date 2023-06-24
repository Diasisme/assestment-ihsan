package migrations

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ihsansolusi.co.id/bcas/service-api-go/models"
	"ihsansolusi.co.id/bcas/service-api-go/pkg/logging"
)


var	log *logging.Logger

func  RunMigration(db *gorm.DB)  {
	err := db.AutoMigrate(&models.Tabungan{}, &models.User{})
	if err != nil {
		remark := "Table already created"
		log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, remark)
		return
	}

	fmt.Printf("Table Migrated")

	
}