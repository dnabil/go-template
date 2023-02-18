package main

import (
	"go-template/database"
	"go-template/handler"
	"go-template/handler/middleware"
	"go-template/repository"
	"go-template/service"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Panicln("Error loading .env file")
	}

	// init database
	db, err := database.InitMySQL()
	if err != nil {
		log.Panicln(err)
	}

	// database migration
	if err = database.MigrateMySQL(db); err != nil {
		log.Panicln(err)
	}
	
	/* architecture: handler -> service/usecase -> repository */
	// repositories
	userRepo := repository.NewUserRepository(db)

	// services
	userService := service.NewUserService(&userRepo)

	// handlers
	userHandler := handler.NewUserHandler(&userService)


	// setup gin
	if (os.Getenv("GIN_MODE") == "release"){
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()
	app.Use(middleware.Timeout(15 * time.Second))
	// app.Use(middleware.PrintBody())
	// app.Use(middleware.PrintHeader())
	
	// app routes
	userHandler.Route(app)
	
	app.Run()
}