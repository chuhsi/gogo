package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	结构体数据和json数据互转
*/

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	// 解析json数据到结构体
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})
	// 结构体转化成json数据
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := &User{
			Firstname: "john",
			Lastname:  "doe",
			Age:       25,
		}
		json.NewEncoder(w).Encode(peter)
	})
	http.ListenAndServe(":9090", nil)
}
