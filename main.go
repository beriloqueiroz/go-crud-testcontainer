package main

import (
	"context"
	"log"

	_ "github.com/crudtest/co-crud-testcontainer/docs"
	"github.com/crudtest/co-crud-testcontainer/src/configuration/database/mongodb"
	"github.com/crudtest/co-crud-testcontainer/src/configuration/logger"
	"github.com/crudtest/co-crud-testcontainer/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Meu Primeiro CRUD em Go | crudtest
// @version 1.0
// @description API for crud operations on users
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func main() {
	logger.Info("About to start user application")

	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
