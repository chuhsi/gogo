package controller

import (
	"gogogo/bookstore/model"
	"gogogo/bookstore/repository"
	"gogogo/bookstore/utils"
	"html/template"
	"net/http"
	"time"
)

func CheckOut(w http.ResponseWriter, r *http.Request) {
	_, session := repository.IsLogin(r)
	userId := session.UserId
	cart, _ := repository.GetCartByUserId(userId)
	orderId := utils.CreateUUID()
	timeStr := time.Now().Format("2006-01-02 12:12:12")
	order := &model.Order{
		OrderId:     orderId,
		CreateTime:  timeStr,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserId:      userId,
	}
	repository.AddOrder(order)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderId: orderId,
		}
		repository.AddOrderItem(orderItem)
		book := v.Book
		book.Sales = book.Sales + v.Count
		book.Stock = book.Stock - v.Count
		repository.UpdateBook(book)
	}
	repository.DelCartByCartId(cart.CartId)
	session.OrderId = orderId
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, session)
}

func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	orderItems, _ := repository.GetOrderItemsByOrderId(orderId)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, _ := repository.GetOrders()
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

func GetMyOrder(w http.ResponseWriter, r *http.Request) {
	_, session := repository.IsLogin(r)
	userId := session.UserId
	orders, _ := repository.GetMyOrder(userId)
	session.Orders = orders
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, session)
}

func SendOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	repository.UpdateOrderState(orderId, 1)
	GetOrders(w, r)
}

func TakeOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	repository.UpdateOrderState(orderId, 2)
	GetMyOrder(w, r)
}
