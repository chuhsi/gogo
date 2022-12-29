package repository

import (
	"gogogo/bookstore/model"
	"gogogo/bookstore/utils"
)

func CheckUsernameAndPassword(username, password string) (*model.User, error) {
	sql := "select id, username, password, email from users where username = ? and password = ?"
	row := utils.DB.QueryRow(sql, username, password)
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func CheckUsername(username string) (*model.User, error) {
	sql := "select id, username, password, email, from users where username = ?"
	row := utils.DB.QueryRow(sql, username)
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func SaveUser(username, password, email string) error {
	sql := "insert into users (username, password, email) values (?, ?, ?)"
	_, err := utils.DB.Exec(sql, username, password, email)
	if err != nil {
		return err
	}
	return nil
}