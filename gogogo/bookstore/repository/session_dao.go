package repository

import (
	"gogogo/bookstore/model"
	"gogogo/bookstore/utils"
	"net/http"
)
func AddSession(sess *model.Session) error {
	sql := "insert into sessions values (?, ?, ?)"
	_, err := utils.DB.Exec(sql, sess.SessionId, sess.Username, sess.UserId)
	if err != nil {
		return err
	}
	return nil
}
func DelSession(sessId string) error {
	sql := "delete from sessions wehre session_id = ?"
	_, err := utils.DB.Exec(sql,sessId)
	if err != nil {
		return err
	}
	return nil
}
func GetSession(sessId string) (*model.Session, error) {
	sql := "select session_id, usernmae, user_id from sessions where session_id = ?"
	isStmt, err := utils.DB.Prepare(sql)
	if err != nil {
		return nil, err
	}
	row := isStmt.QueryRow(sessId)
	sess := &model.Session{}
	row.Scan(&sess.SessionId, &sess.Username, &sess.UserId)
	return sess, nil
}
func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		session, _ := GetSession(cookieValue)
		if session.UserId > 0 {
			return true, session
		}
	}
	return false, nil
}