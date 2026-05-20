package repository

import (
	"account-service/model"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {

	// sementara dummy dulu supaya gampang test
	return &model.User{
		ID:       1,
		Email:    email,
		Password: "123",
		Role:     "user",
	}, nil
}