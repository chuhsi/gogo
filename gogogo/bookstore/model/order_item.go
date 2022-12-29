package model

type OrderItem struct {
	OrderItemId int
	Count int
	Amount float64
	Title string
	Author string
	Price float64
	ImgPath string
	OrderId string
}