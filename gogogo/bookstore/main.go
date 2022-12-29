package main

import (
	"gogogo/bookstore/controller"
	"log"
	"net/http"
)

func main() {
	// sql := "select id, username, password, email from users where username = ? and password = ?"
	// r := utils.DB.QueryRow(sql, "lisi", "123123")
	// user := model.User{}
	// r.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	// fmt.Println(user)
	// 处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/", controller.GetPageBookByPrice)
	http.HandleFunc("/main", controller.GetPageBookByPrice)
	http.HandleFunc("/checkUserName", controller.CheckUsername)
	// 登陆模块
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/regist", controller.Register)
	// 图书模块
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBookByPrice)
	http.HandleFunc("/deleteBook", controller.DelBook)
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	http.HandleFunc("/updateOrandBook", controller.UpdateOrAddBook)
	// 购物车模块
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	http.HandleFunc("/deleteCart", controller.DelCart)
	http.HandleFunc("/delelteCartItem", controller.DelCartItem)
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	http.HandleFunc("/checkout", controller.CheckOut)
	// 订单模块
	http.HandleFunc("/getOrders", controller.GetOrders)
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	http.HandleFunc("/getMyOrder", controller.GetMyOrder)
	http.HandleFunc("/sendOrder", controller.SendOrder)
	http.HandleFunc("/takeOrder", controller.TakeOrder)
	log.Println("程序启动成功 ...")
	http.ListenAndServe(":9090", nil)
}
