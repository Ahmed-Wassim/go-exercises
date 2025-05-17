package main

import (
	"encoding/json"
	"net/http"
	"text/template"
)

type todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{.ID}}</h1>
<h3>{{printf "User ID: %d" .UserID}}</h3>
<h3>{{printf "Title: %s" .Title}}</h3>
<h3>{{printf "Completed: %t" .Completed}}</h3>
`

func handler(w http.ResponseWriter, r *http.Request) {
	const base = "https://jsonplaceholder.typicode.com/"

	resp, err := http.Get(base + r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var item todo

	err = json.NewDecoder(resp.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temp := template.New("todo")
	temp.Parse(form)
	temp.Execute(w, item)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
