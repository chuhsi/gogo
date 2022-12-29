package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt time.Time
}

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// single insert
	// user := &User{
	// 	Id: 0,
	// 	Username: "lisi",
	// 	Password: "123",
	// 	CreatedAt: time.Now(),
	// }
	// batch insert

	// users := []User{{Id: 0,
	// 	Username:  "wangwu",
	// 	Password:  "123",
	// 	CreatedAt: time.Now()}, {Id: 1,
	// 	Username:  "lily",
	// 	Password:  "123",
	// 	CreatedAt: time.Now()}}
	// d := db.Select("username", "password", "created_at")
	// d := db.Table("users").Create(map[string]interface{}{
	// 	"username":  "max",
	// 	"password":  "123",
	// 	"created_at": time.Now(),
	// })

	// for _, user := range users {
	// 	fmt.Println("Username: ", user.Username)
	// }

	var user []User

	// d := db.Find(&user)

	// tx := db.Model(&user).Where("password",123).Update("password",111)
	// if i := tx.RowsAffected; i <= 0 {
	// 	fmt.Println(i)
	// 	log.Fatal(tx.Error)
	// 	return
	// }
	tx := db.Where("username","job").Delete(&user)
	if i := tx.RowsAffected; i <= 0 {
		fmt.Println(i)
		log.Fatal(tx.Error)
		return
	}
	// fmt.Println(user)
}
