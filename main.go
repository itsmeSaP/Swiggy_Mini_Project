package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/itsmeSaP/Swiggy_Mini_Project/pkg/database"
	model "github.com/itsmeSaP/Swiggy_Mini_Project/pkg/models"
)

func main() {
	//Connecting to DATABASE
	db, err := database.ConnectDB()
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println("Connected to DB")

	//RUNNING AUTOMIGRATIONS
	err = db.AutoMigrate(&model.User{}, &model.City{}, &model.Address{}, &model.Restaurent{}, &model.Menu{}, &model.Order{}, &model.FoodCategory{}, &model.ItemsOrdered{}, &model.Payment{})
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println("Migrations Done")

	//creating the fiber app
	app := fiber.New()
	//Middleware Initialization
	app.Use(cors.New())
	app.Use(logger.New())

	log.Fatal(app.Listen(":3000"))
}
