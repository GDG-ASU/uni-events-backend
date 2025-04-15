package main

import (
	"uni-events-backend/config"
	"uni-events-backend/internal/api/user"
	"uni-events-backend/internal/middlewares"
	"uni-events-backend/internal/repositories"
	"uni-events-backend/internal/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := config.InitDB()
	
	// db.AutoMigrate(&models.User{}) 
	
	userRepo := repositories.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := user.NewHandler(userService)

	// Group routes
	apiGroup := e.Group("/api/v1")
	userGroup := apiGroup.Group("/users")
	userGroup.Use(middlewares.ClerkAuthMiddleware)
	userGroup.GET("/me", userHandler.Me)

	e.Logger.Fatal(e.Start(":8080"))
}