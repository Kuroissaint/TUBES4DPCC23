package repository

import "account-service/model"

type UserRepository interface {
	GetByEmail(email string) (*model.User, error)
}