package controller

import (
	"encoding/json"
	"gogogo/bookstore/model"
	"gogogo/bookstore/repository"
	"gogogo/bookstore/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	flag, session := repository.IsLogin(r)
	if flag {
		bookId := r.FormValue("bookId")
		book, _ := repository.GetBookById(bookId)
		userId := session.UserId
		cart, _ := repository.GetCartByUserId(userId)
		if cart != nil {
			cartItem, _ := repository.GetCartItemByBookIdAndCartId(bookId, cart.CartId)
			if cartItem != nil {
				cts := cart.CartItems
				for _, v := range cts {
					log.Println("当前购物项中是否有Book", v)
					log.Println("查到的Book是", cartItem.Book)
					if v.Book.Id == cartItem.Book.Id {
						v.Count = v.Count + 1
						repository.UpdateBookCount(v)
					}
				}
			} else {
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartId: cart.CartId,
				}
				cart.CartItems = append(cart.CartItems, cartItem)
				repository.AddCartItem(cartItem)
			}
			repository.UpdateCart(cart)
		} else {
			cartId := utils.CreateUUID()
			cart := &model.Cart{
				CartId: cartId,
				UserId: userId,
			}
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartId: cartId,
			}
			cartItems = append(cartItems, cartItem)
			cart.CartItems = cartItems
			repository.AddCart(cart)
		}
		w.Write([]byte("你刚刚将" + book.Title + "添加到了购物车"))
	} else {
		w.Write([]byte("请先登录"))
	}
}

func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := repository.IsLogin(r)
	uesrId := session.UserId
	cart, _ := repository.GetCartByUserId(uesrId)
	if cart != nil {
		session.Cart = cart
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	} else {
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	}
}

func DelCart(w http.ResponseWriter, r *http.Request) {
	cartId := r.FormValue("cartId")
	repository.DelCartByCartId(cartId)
	GetCartInfo(w, r)
}

func DelCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.FormValue("cartItemId")
	iCartItemId, _ := strconv.ParseInt(cartItemId, 10, 64)
	_, session := repository.IsLogin(r)
	userId := session.UserId
	cart, _ := repository.GetCartByUserId(userId)
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if v.CartItemId == int(iCartItemId) {
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			cart.CartItems = cartItems
			repository.DelCartItemsById(cartItemId)
		}
	}
	repository.UpdateCart(cart)
	GetCartInfo(w, r)
}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.FormValue("cartItemId")
	bookCount := r.FormValue("bookCount")
	iCartItemId, _ := strconv.ParseInt(cartItemId, 10, 64)
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	_, session := repository.IsLogin(r)
	userId := session.UserId
	cart, _ := repository.GetCartByUserId(userId)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if v.CartItemId == int(iCartItemId) {
			v.Count = int(iBookCount)
			repository.UpdateBookCount(v)
		}
	}
	repository.UpdateCart(cart)
	cart, _ = repository.GetCartByUserId(userId)
	totalCount := cart.TotalCount
	totalAmount := cart.TotalAmount
	var amount float64
	cis := cart.CartItems
	for _, v := range cis {
		if iCartItemId == int64(v.CartItemId) {
			amount = v.Amount
		}
	}
	data := &model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	json, _ := json.Marshal(data)
	w.Write(json)
}
