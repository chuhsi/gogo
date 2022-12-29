package repository

import (
	"gogogo/bookstore/model"
	"gogogo/bookstore/utils"
)

func AddCartItem(cartItem *model.CartItem) error {
	sql := "insert into cart_items (count, amount, book_id, cart_id) values (?, ?, ?, ?)"
	_, err := utils.DB.Exec(sql, cartItem.Count, cartItem.Amount, cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		return err
	}
	return nil
}

func GetCartItemByBookIdAndCartId(bookId, cartId string) (*model.CartItem, error) {
	sql := "select id, count, amount, cart_id, from cart_items where book_id = ? and cart_id = ?"
	row := utils.DB.QueryRow(sql)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartId)
	if err != nil {
		return nil, err
	}
	book, _ := GetBookById(bookId)
	cartItem.Book = book
	return cartItem, nil
}

func UpdateBookCount(cartItem *model.CartItem) error {
	sql := "update cart_items set count = ?, amount = ?, where book_id = ? and cart_id = ?"
	_, err := utils.DB.Exec(sql, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		return err
	}
	return nil
}

func GetCartItemByCartId(cartId string) ([]*model.CartItem, error) {
	sql := "select id, count, amount, book_id, cart_id from cart_items where cart_id = ?"
	rows, err := utils.DB.Query(sql, cartId)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		var bookId string
		cartItem := &model.CartItem{}
		err := rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartId)
		if err != nil {
			return nil, err
		}
		book, _ := GetBookById(bookId)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

func DelCartItemsByCartId(cartId string) error {
	sql := "delete from cart_items where cart_id = ?"
	_, err := utils.DB.Exec(sql, cartId)
	if err != nil {
		return err
	}
	return nil
}

func DelCartItemsById(cartItemId string) error {
	sql := "delete from cart_items where id = ?"
	_, err := utils.DB.Exec(sql, cartItemId)
	if err != nil {
		return err
	}
	return nil
}
