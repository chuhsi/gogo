package model

type Order struct {
	OrderId string
	CreateTime string
	TotalCount int
	TotalAmount float64
	State int
	UserId int
}
func (order *Order) NoSend() bool {
	return order.State == 0
}
func (order *Order) SendComplete() bool {
	return order.State == 1
}
func (order *Order) Complete() bool {
	return order.State == 2
}