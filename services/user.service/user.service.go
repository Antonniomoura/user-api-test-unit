package user_service

import (
	"awesomeProject/models"
	"awesomeProject/repositories/user.repository"
)

func Create(user models.User) error {
	err := user_repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func Read() (models.Users, error) {
	users, err := user_repository.Read()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func Update(user models.User, userId string) error {
	err := user_repository.Update(user, userId)
	if err != nil {
		return err
	}
	return nil
}

func Delete(userId string) error {
	err := user_repository.Delete(userId)

	if err != nil {
		return err
	}
	return nil
}
