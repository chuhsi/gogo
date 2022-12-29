package model

type Page struct {
	Books []*Book
	PageNo int
	PageSize int
	TotalPageNo int
	TotalRecord int
	MinPrice string
	MaxPrice string
	IsLogin bool
	Username string
}
func (page *Page) HasPre() bool {
	return page.PageNo > 1
}
func (page *Page) HasNext() bool {
	return page.PageNo < page.TotalPageNo
}
func (page *Page) GetPrePageNo() int {
	if page.HasPre() {
		return page.PageNo - 1
	}
	return 1
}
func (page *Page) GetNextPageNo() int {
	if page.HasPre() {
		return page.PageNo + 1
	}
	return page.TotalPageNo
}