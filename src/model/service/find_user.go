package service

import (
	"github.com/crudtest/co-crud-testcontainer/src/configuration/logger"
	"github.com/crudtest/co-crud-testcontainer/src/configuration/rest_err"
	"github.com/crudtest/co-crud-testcontainer/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID services.",
		zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services.",
		zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services.",
		zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
