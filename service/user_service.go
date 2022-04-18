package service

import (
	"user-service/db"
	"user-service/model"
	"user-service/repository"
)

func GetUser(id int) (*model.UserModel, error) {
	userRepository := repository.UserNeo4jRepository{
		Drive: db.GetNeo4j(),
	}

	user, err := userRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
