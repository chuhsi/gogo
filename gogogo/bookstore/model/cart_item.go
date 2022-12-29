package model

type CartItem struct {
	CartItemId int
	Book *Book
	Count int
	Amount float64
	CartId string
}

func (cartItem *CartItem) GetAmount() float64 {
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}