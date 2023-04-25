package api

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	authAPI "github.com/dancondo/users-api/api/auth-api"
	healthAPI "github.com/dancondo/users-api/api/health-api"
	"github.com/dancondo/users-api/common"
	_ "github.com/dancondo/users-api/docs/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title User API
// @version 1.0
// @description User API

// StartHTTP server
func StartHTTP() error {
	common.LoadEnv()

	srv := CreateRouter()
	port := common.GetEnv("HTTP_PORT")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c //nolint:gosimple
		fmt.Println("Gracefully shutting down...")
		_ = srv.Shutdown()
	}()

	if err := srv.Listen(fmt.Sprintf(":%s", port)); err != nil {
		common.Log.Panic(err)
	}

	return nil
}

func CreateRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Duration(common.GetEnvInt("HTTP_SERVER_READ_TIMEOUT_SECONDS")) * time.Second,
		WriteTimeout: time.Duration(common.GetEnvInt("HTTP_SERVER_WRITE_TIMEOUT_SECONDS")) * time.Second,
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(compress.New())

	router := app.Group("/api")

	authRouter := router.Group(authAPI.HandlerPath)
	healthRouter := router.Group(healthAPI.HandlerPath)

	authAPI.RegisterRoutes(authRouter)
	healthAPI.RegisterRoutes(healthRouter)

	app.Get("/docs/swagger/*", fiberSwagger.WrapHandler)

	return app
}
