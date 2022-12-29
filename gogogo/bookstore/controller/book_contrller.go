package controller

import (
	"gogogo/bookstore/model"
	"gogogo/bookstore/repository"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func GetPageBookByPrice(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		page, _ = repository.GetPageBooks(pageNo)
		log.Println(page)
	} else {
		page, _ = repository.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
		log.Println(page)
	}
	// 判断是否登陆
	flag, session := repository.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = session.Username
	}
	// 解析模板文件
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := repository.GetPageBooks(pageNo)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

func DelBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	repository.DelBook(bookId)
	GetPageBooks(w, r)
}

func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	book, _ := repository.GetBookById(bookId)
	if book.Id > 0 {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
}

func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	iPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	iBookId, _ := strconv.ParseInt(bookId, 10, 0)

	book := &model.Book{
		Id:      int(iBookId),
		Title:   title,
		Author:  author,
		Price:   iPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	if book.Id > 0 {
		repository.UpdateBook(book)
	} else {
		repository.AddBook(book)
	}
	GetPageBooks(w, r)
}
