package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func getMySql() *sql.DB {

	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/studentinfo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func api(w http.ResponseWriter, r *http.Request) {
	db = getMySql()
	if r.URL.Path != "/" {
		http.Error(w, "404 PAGE NOT FOUND ", http.StatusNotFound)

		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "forms.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err:%v", err)
			return
		}
		fmt.Fprintf(w, "Post  from website r.postfrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		id := r.FormValue("id")

		fmt.Fprintf(w, "Name =%s\n", name)
		fmt.Fprintf(w, "Id =%s\n", id)

	default:
		fmt.Fprintf(w, "Only get and post")
	}
}

func main() {
	http.HandleFunc("/", api)
	fmt.Printf("Server Got testing \n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
