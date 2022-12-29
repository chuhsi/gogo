package main

import (
	"database/sql"
	// "fmt"
	"log"
	// "time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// create a new talbe
	// {
	// 	query := `
	// 		create table users (
	// 			id int(11) auto_increment primary key,
	// 			username varchar(20) not null,
	// 			password varchar(20) not null,
	// 			created_at timestamp
	// 		);
	// 	`
	// 	if _, err := db.Exec(query); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// // insert a new user
	// {
	// 	username := "job"
	// 	password := "123"
	// 	createdAt := time.Now()

	// 	rus, err := db.Exec(`insert into users (username, password, created_at) values (?,?,?)`, username, password, createdAt)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	i, err := rus.LastInsertId()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(i)
	// }
	// // query a single user
	// {
	// 	var (
	// 		id int
	// 		username string
	// 		password string
	// 		createdAt time.Time
	// 	)
	// 	query := "select id, username, password, created_at from users where id = ?"
	// 	if err := db.QueryRow(query, 4).Scan(&id, &username, &password, &createdAt); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(id, username, password, createdAt)
	// }
	// // query all users
	// {
	// 	type User struct {
	// 		id int
	// 		username string
	// 		password string
	// 		createdAt time.Time
	// 	}
	// 	query := "select id, username, password, created_at from users"
	// 	rows, err := db.Query(query)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer rows.Close()

	// 	var users []User
	// 	for rows.Next() {
	// 		var u User
	// 		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		users = append(users, u)
	// 	}
	// 	if err := rows.Err(); err != nil{
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%#v\n",users)
	// }
	// // deleted table user
	{
		_,err := db.Exec(`delete from users where id = ?`,1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
