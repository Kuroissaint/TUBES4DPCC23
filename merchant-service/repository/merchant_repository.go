package repository

import "merchant-service/model"

type MerchantRepository interface {
	GetByNamaToko(namaToko string) (*model.Merchant, error)
	RegisterMerchant(req model.MerchantRegisterRequest, hashedPassword string) (*model.Merchant, error)
	GetMerchantByID(id int) (*model.Merchant, error)
	AddMenuItem(merchantID int, req model.MenuItemRequest) (*model.MenuItem, error)
	GetMenuByMerchantID(merchantID int) ([]model.MenuItem, error)
	UpdateMenuItem(id int, req model.MenuItemRequest) (*model.MenuItem, error)
	DeleteMenuItem(id int) error
}
