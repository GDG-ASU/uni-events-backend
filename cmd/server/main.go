package main

import (
	"uni-events-backend/config"
	"uni-events-backend/internal/api/club"
	"uni-events-backend/internal/api/user"
	"uni-events-backend/internal/middlewares"
	"uni-events-backend/internal/models"
	"uni-events-backend/internal/repositories"
	"uni-events-backend/internal/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := config.InitDB()
	
   db.AutoMigrate(&models.User{},&models.Club{},&models.ClubOwner{}) 

	userRepo := repositories.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := user.NewHandler(userService)


	clubRepo := repositories.NewClubRepository(db)
	clubService := service.NewClubService(clubRepo)
	clubHandler := club.NewHandler(clubService, userService)


	// Group routes
	apiGroup := e.Group("/api/v1")
	userGroup := apiGroup.Group("/users")
	userGroup.Use(middlewares.ClerkAuthMiddleware)
	userGroup.GET("/getme", userHandler.GetMe)

	clubGroup := apiGroup.Group("/clubs")
	clubGroup.Use(middlewares.ClerkAuthMiddleware)
	clubGroup.POST("/create-club",clubHandler.CreateClub)
	clubGroup.PATCH("/clubs/:id", clubHandler.UpdateClub)


	e.Logger.Fatal(e.Start(":8080"))
}