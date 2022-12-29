package repository

import (
	"gogogo/bookstore/model"
	"gogogo/bookstore/utils"
)

func AddOrderItem(orderItem *model.OrderItem) error {
	sql := "insert into order_items (count, amount, title, price, img_path, order_id) values (?, ?, ?, ?, ?, ?)"
	_, err := utils.DB.Exec(sql, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Price, orderItem.ImgPath, orderItem.OrderId)
	if err != nil {
		return err
	}
	return nil
}

func GetOrderItemsByOrderId(orderId string) ([]*model.OrderItem, error) {
	sql := "select id, count, amount, title, author, price, img_path, order_id from order_items where order_id = ?"
	rows, err := utils.DB.Query(sql, orderId)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemId, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderId)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}

