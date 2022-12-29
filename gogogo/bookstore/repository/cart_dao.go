package repository

import (
	"gogogo/bookstore/model"
	"gogogo/bookstore/utils"
)

func AddCart(cart *model.Cart) error {
	sql := "insert into carts (id, total_count, total_amount, user_id) values (?, ?, ?, ?)"
	_, err := utils.DB.Exec(sql, cart.CartId, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserId)
	if err != nil {
		return err
	}
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		AddCartItem(cartItem)
	}
	return nil
}

func GetCartByUserId(userId int) (*model.Cart, error) {
	sql := "select id, total_count, total_amount, user_id from carts where user_id = ?"
	row := utils.DB.QueryRow(sql, userId)
	cart := &model.Cart{}
	err := row.Scan(&cart.CartId, &cart.TotalCount, &cart.TotalAmount, &cart.UserId)
	if err != nil {
		return nil, err
	}
	cartItems, _ := GetCartItemByCartId(cart.CartId)
	cart.CartItems = cartItems
	return cart, nil
}

func UpdateCart(cart *model.Cart) error {
	sql := "udpate carts set total_count = ?, total_amount = ? where id = ?"
	_, err := utils.DB.Exec(sql, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartId)
	if err != nil {
		return err
	}
	return nil
}

func DelCartByCartId(cartId string) error {
	err := DelCartItemsByCartId(cartId)
	if err != nil {
		return err
	}
	sql := "delete from carts wherre id = ?"
	_, err = utils.DB.Exec(sql, cartId)
	if err != nil {
		return err
	}
	return nil
}
