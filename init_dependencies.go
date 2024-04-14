package main

import (
	"github.com/crudtest/co-crud-testcontainer/src/controller"
	"github.com/crudtest/co-crud-testcontainer/src/model/repository"
	"github.com/crudtest/co-crud-testcontainer/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
