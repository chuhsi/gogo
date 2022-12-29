package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

func main() {
	temp,_ := template.ParseFiles("index.html")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := &TodoPageData{
			PageTitle: "my todo list",
			Todos: []Todo{
				{Title: "task1", Done: false},
				{Title: "task2", Done: true},
				{Title: "task3", Done: true},
			},
		}
		temp.Execute(w, data)
	})
	http.ListenAndServe(":9090", nil)
}