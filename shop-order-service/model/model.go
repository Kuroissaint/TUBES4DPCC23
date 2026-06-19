package model

type ShoppingCart struct {
	OrderID         string
	UserID          string
	MerchantID      string
	Items           []string
	TotalPrice      float64
	DeliveryAddress string
	PaymentStatus   string
	Status          string
}

func (c *ShoppingCart) AddToCart(item string) {
	c.Items = append(c.Items, item)
}