package repository

import "merchant-service/model"

type MerchantRepository interface {
	GetByName(name string) (*model.Merchant, error)
}