package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/khunfloat/sgcu-borrow-backend/handler"
	"github.com/khunfloat/sgcu-borrow-backend/logs"
	"github.com/khunfloat/sgcu-borrow-backend/repository"
	"github.com/khunfloat/sgcu-borrow-backend/service"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	// Load config
	initConfig()
	// Set timezone
	initTimeZone()
	// Connect database
	db := initDatabase()
	// Init fiber app
	app := initApp()
	app.Use(cors.New())
	api := app.Group("/api")

	userRepository := repository.NewUserRepositoryDB(db)
	userAuthService := service.NewUserAuthService(userRepository)


	staffRepository := repository.NewStaffRepositoryDB(db)
	staffAuthService := service.NewStaffAuthService(staffRepository)

	authHandler := handler.NewAuthHandler(userAuthService, staffAuthService)

	itemRepository := repository.NewItemRepositoryDB(db)
	itemService := service.NewItemService(itemRepository)
	itemHandler := handler.NewItemHandler(itemService)

	// public api
	api.Post("/signup", authHandler.UserSignUp)
	api.Post("/signin", authHandler.UserSignIn)
	api.Post("/staff/signup", authHandler.StaffSignUp)
	api.Post("/staff/signin", authHandler.StaffSignIn)

	// user api
	api.Use(authHandler.AuthorizationRequired())

	api.Get("/items", itemHandler.GetItems)
	api.Get("/item/:item_id", itemHandler.GetItem)

	// staff api
	api.Use(authHandler.IsStaff)

	api.Post("/item/create", itemHandler.CreateItem)
	api.Put("/item/update", itemHandler.UpdateItem)
	api.Delete("/item/:item_id", itemHandler.DeleteItem)

	// admin api
	api.Use(authHandler.IsAdmin)

	// Start server
	logs.Info("CUBS coin service started at port " + viper.GetString("app.port"))
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"))

	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
        TranslateError: true,
	})
	if err != nil {
		panic(err)
	}

	return db
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initApp() *fiber.App {
    return fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               viper.GetBool("app.prefork"),
		AppName:               viper.GetString("app.name"),
	})
}