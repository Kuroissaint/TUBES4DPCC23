package repository

import "account-service/model"

type UserRepository interface {
	GetByEmail(email string) (*model.User, error)
	Register(user model.User) (*model.User, error)
	RegisterDriver(user model.User, driver model.DriverProfile) (*model.User, error)
	RegisterCustomer(user model.User, alamat string) (*model.User, error)
}
