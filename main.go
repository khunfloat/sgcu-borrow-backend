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

	userAPI := api.Group("")
	staffAPI := api.Group("")

	userRepository := repository.NewUserRepositoryDB(db)
	userAuthService := service.NewUserAuthService(userRepository)
	userAuthHandler := handler.NewUserAuthHandler(userAuthService)

	staffRepository := repository.NewStaffRepositoryDB(db)
	staffAuthService := service.NewStaffAuthService(staffRepository)
	staffAuthHandler := handler.NewStaffAuthHandler(staffAuthService)

	itemRepository := repository.NewItemRepositoryDB(db)
	itemService := service.NewItemService(itemRepository)
	itemHandler := handler.NewItemHandler(itemService)

	// public api
	userAPI.Post("/signup", userAuthHandler.SignUp)
	userAPI.Post("/signin", userAuthHandler.SignIn)
	staffAPI.Post("/staff/signup", staffAuthHandler.SignUp)
	staffAPI.Post("/staff/signin", staffAuthHandler.SignIn)

	// user api
	userAPI.Get("/items", itemHandler.GetItems)
	userAPI.Get("/item/:item_id", itemHandler.GetItem)

	// staff api
	staffAPI.Use(staffAuthHandler.AuthorizationRequired())
	staffAPI.Use(staffAuthHandler.IsStaff)

	staffAPI.Post("/item/create", itemHandler.CreateItem)
	staffAPI.Put("/item/update", itemHandler.UpdateItem)
	staffAPI.Delete("/item/:item_id", itemHandler.DeleteItem)

	// admin api

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