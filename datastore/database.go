package datastore

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ihsansolusi.co.id/bcas/service-api-go/pkg/logging"
)

type TabunganDatabase struct {
	DB  *gorm.DB
	log logging.Logger
}

func (t *TabunganDatabase) DatabaseInit(host, user, password, dbname, timezone, port string) (err error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, dbname, port, timezone)
	t.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Connected to databse")
	return err
}

// migrate -path datastore/migration/ -database "postgresql://user:password@localhost:5436/assestment-ihsan?sslmode=disable" -verbose up
