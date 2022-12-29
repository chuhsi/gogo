package repository

import (
	"gogogo/bookstore/model"
	"gogogo/bookstore/utils"
)

func AddOrder(order *model.Order) error {
	sql := "insert into orders (id, create_time, total_count, total_amount, state, user_id) values (?, ?, ?, ?, ?, ?)"
	_, err := utils.DB.Exec(sql, order.OrderId, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserId)
	if err != nil {
		return err
	}
	return nil
}

func GetOrders() ([]*model.Order, error) {
	sql := "select id, create_time, total_count, total_amount, state, user_id from orders"
	rows, err := utils.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		orders = append(orders, order)
	}
	return orders, nil
}

func GetMyOrder(userId int) ([]*model.Order, error) {
	sql := "select id, create_time, total_count, total_amount, state, user_id from orders where user_id = ?"
	rows, err := utils.DB.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		orders = append(orders, order)
	}
	return orders, nil
}

func UpdateOrderState(orderId string, state int) error {
	sql := "udpate orders set state = ? where id = ?"
	_, err := utils.DB.Exec(sql, state, orderId)
	if err != nil {
		return err
	}
	return nil
}
