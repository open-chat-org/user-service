package service

import (
	"user-service/model"
	"user-service/repository"
)

type UserNeo4jService struct {
	UserRepository repository.UserNeo4jRepository
}

func (u *UserNeo4jService) GetUser(id int) (*model.UserModel, error) {
	user, err := u.UserRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserNeo4jService) DeleteUser(id int) error {
	err := u.UserRepository.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserNeo4jService) UpdateAvatar(id int, avatar string) error {
	err := u.UserRepository.UpdateAvatarById(id, avatar)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserNeo4jService) UpdateProfile(id int, avatar string) error {
	err := u.UserRepository.UpdateAvatarById(id, avatar)
	if err != nil {
		return err
	}

	return nil
}
