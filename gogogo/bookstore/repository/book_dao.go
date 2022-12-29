package repository

import (
	"gogogo/bookstore/model"
	"gogogo/bookstore/utils"
	"strconv"
)

// 获取所有图书
func GetBooks() ([]*model.Book, error) {
	sql := "select id, title, author, price, sales, stock, img_path from books"
	rows, err := utils.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

// 添加一本图书
func AddBook(b *model.Book) error {
	sql := "insert into books (title, author, price, sales, stock, img_path) values (?, ?, ?, ?, ?, ?)"
	_, err := utils.DB.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// 根据图书id删除图书
func DelBook(bookId string) error {
	sql := "delete from books where id = ?"
	_, err := utils.DB.Exec(sql, bookId)
	if err != nil {
		return err
	}
	return nil
}

// 根据图书id查询一本图书
func GetBookById(bookId string) (*model.Book, error) {
	sql := "select id, title, author, price, sales, stock, img_path from books where id = ?"
	row := utils.DB.QueryRow(sql, bookId)
	book := &model.Book{}
	row.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

// 根据图书id跟新图书信息
func UpdateBook(b *model.Book) error {
	sql := "update books set title = ?, author = ?, price = ?, sales = ?, stock = ? where id = ?"
	_, err := utils.DB.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.Id)
	if err != nil {
		return err
	}
	return nil
}

// 获取分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	sql1 := "select count(*) from books"
	var totalRecord int
	row := utils.DB.QueryRow(sql1)
	row.Scan(&totalRecord)
	var pageSize int = 4
	var totalPageNo int
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	sql2 := "select id, title, author, price, sales, stock, img_path from books limit ?, ?"
	rows, err := utils.DB.Query(sql2, (iPageNo-1)*int64(pageSize), pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	page := &model.Page{
		Books:       books,
		PageNo:      int(iPageNo),
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

// 获取带分页和价格范围的图书信息
func GetPageBooksByPrice(pageNo, minPrice, maxPrice string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	sql1 := "select count(*) from books where price between ? and ?"
	var totalRecord int
	row := utils.DB.QueryRow(sql1, minPrice, maxPrice)
	row.Scan(&totalRecord)
	var pageSize int = 4
	var totalPageNo int
	if totalPageNo%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord / pageSize + 1
	}
	sql2 := "select id, title, author, price, sales, stock, img_path from books where price between ? and ? limit ?, ?"
	rows, err := utils.DB.Query(sql2, minPrice, maxPrice, (iPageNo-1)*int64(pageSize), pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	page := &model.Page{
		Books:       books,
		PageNo:      int(iPageNo),
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}
