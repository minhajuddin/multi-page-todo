package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Login struct {
	ID    int64  `db:"id"`
	Email string `db:"email"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error
	w.Header().Add("content-type", "text/html")
	b := strings.Builder{}
	logins := []Login{}
	err = db.Select(&logins, "SELECT id, email From logins")
	log.Printf("ERROR: %v\n", err)
	b.WriteString("<ul>")
	for i, l := range logins {
		fmt.Fprintf(w, "<li>%d %d %s</li>", i+1, l.ID, l.Email)
	}
	b.WriteString("</ul>")
	fmt.Fprintf(w, `<doctype html>
		%s
		<form method=post action="/todo/">
		<input type=text name=title />
		<button type=submit>Save</button>
		</form>
		`, b.String())
}

func CreateTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	log.Printf("ERROR Parsing %v\n", err)
	log.Printf("%v", r.PostForm)
	if title, ok := r.PostForm["title"]; ok {
		log.Printf("Creating task '%s'\n", title[0])
		tx := db.MustBegin()
		tx.MustExec("INSERT INTO logins (email) VALUES($1)", title[0])
		tx.Commit()
	}
	w.Header().Add("location", "/?info=task-created")
	w.WriteHeader(302)
}

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Connect("postgres", "user=postgres password=postgres dbname=minhajuddin sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	logins := []Login{}
	err = db.Select(&logins, "SELECT id, email From logins")
	log.Printf("ERROR: %v\n", err)
	fmt.Println("LOGINS", logins)

	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/todo", CreateTodo)

	log.Fatal(http.ListenAndServe(":8080", router))
}
