package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"ihsansolusi.co.id/bcas/service-api-go/datastore"
	"ihsansolusi.co.id/bcas/service-api-go/datastore/migrations"
	"ihsansolusi.co.id/bcas/service-api-go/route"
)



func main()  {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//DB Config
	DB_HOST := os.Getenv("POSTGRES_HOST")
	DB_USER := os.Getenv("POSTGRES_USER")
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB_NAME := os.Getenv("POSTGRES_DB")
	DB_TIMEZONE := os.Getenv("POSTGRES_TIMEZONE")
	DB_PORT := os.Getenv("POSTGRES_DB_EXPOSE_PORT")


	// Init Database
	ds := datastore.TabunganDatabase{}
	err = ds.DatabaseInit(DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_TIMEZONE, DB_PORT)
	if err != nil {
		panic(err)
	}
	// Init Migration
	migrations.RunMigration(ds.DB)

	app := fiber.New()
	
	// Init Route
	route.RouteInit(app)
}