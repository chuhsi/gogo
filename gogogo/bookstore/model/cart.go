package model

type Cart struct {
	CartId string
	CartItems []*CartItem
	TotalCount int
	TotalAmount float64
	UserId int
}

func (cart *Cart) GetTotalCount() int{
	var totalCount int
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}