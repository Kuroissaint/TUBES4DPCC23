package model

type ShoppingCart struct {
	OrderID    string
	UserID     string
	MerchantID string
	Items      []string
	Status     string
}

func (c *ShoppingCart) AddToCart(item string) {
	c.Items = append(c.Items, item)
}
