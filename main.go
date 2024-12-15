package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	_ "github.com/khunfloat/sgcu-borrow-backend/docs"
	"github.com/khunfloat/sgcu-borrow-backend/handler"
	"github.com/khunfloat/sgcu-borrow-backend/logs"
	"github.com/khunfloat/sgcu-borrow-backend/repository"
	"github.com/khunfloat/sgcu-borrow-backend/service"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// @title           SGCU Borrowing System API
// @version         1.0
// @description     This is an example server.
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://example.com/contact
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8000
// @BasePath  /api

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

	docs := app.Group("/docs")
	docs.Get("/*", swagger.HandlerDefault)

	api := app.Group("/api")

	userRepository := repository.NewUserRepositoryDB(db)
	userAuthService := service.NewUserAuthService(userRepository)

	staffRepository := repository.NewStaffRepositoryDB(db)
	staffAuthService := service.NewStaffAuthService(staffRepository)

	authHandler := handler.NewAuthHandler(userAuthService, staffAuthService)

	itemRepository := repository.NewItemRepositoryDB(db)
	itemService := service.NewItemService(itemRepository)
	itemHandler := handler.NewItemHandler(itemService)

	orderRepository := repository.NewOrderRepositoryDB(db)
	borrowRepository := repository.NewBorrowRepositoryDB(db)
	returnRepository := repository.NewReturnRepositoryDB(db)
	lostRepository := repository.NewLostRepositoryDB(db)

	orderService := service.NewOrderService(
		orderRepository,
		itemRepository,
		borrowRepository,
		returnRepository,
		lostRepository,
	)
	orderHandler := handler.NewOrderHandler(orderService)

	// public api
	api.Post("/signup", authHandler.UserSignUp)
	api.Post("/signin", authHandler.UserSignIn)
	api.Post("/staff/signup", authHandler.StaffSignUp)
	api.Post("/staff/signin", authHandler.StaffSignIn)

	api.Get("/orders", orderHandler.GetOrders)
	api.Get("/order/:order_id", orderHandler.GetOrder)
	api.Post("/order", orderHandler.CreateOrder)
	api.Put("/order", orderHandler.UpdateOrder)
	api.Delete("/order/:order_id", orderHandler.DeleteOrder)

	api.Post("/pickup", orderHandler.PickupOrder)

	// user api
	// api.Use(authHandler.AuthorizationRequired())

	

	// staff api
	// api.Use(authHandler.IsStaff)
	api.Get("/items", itemHandler.GetItems)
	api.Get("/item/:item_id", itemHandler.GetItem)
	api.Post("/item", itemHandler.CreateItem)
	api.Put("/item", itemHandler.UpdateItem)
	api.Delete("/item/:item_id", itemHandler.DeleteItem)

	// admin api
	// api.Use(authHandler.IsAdmin)

	// Start server
	logs.Info("SGCU borrowing service started at port " + viper.GetString("app.port"))
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
		Logger:         logger.Default.LogMode(logger.Silent),
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
	logs.Info("load timezone Asia/Bangkok")
}

func initApp() *fiber.App {
	return fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               viper.GetBool("app.prefork"),
		AppName:               viper.GetString("app.name"),
	})
}