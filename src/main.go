package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type PageData struct {
	Statuscode string
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index.html"))

	r.ParseForm()
	var code string
	if r.FormValue("kenteken") != "" {
		if len(r.FormValue("kenteken")) != 8 {
			code = "Geen geldig kenteken"
		} else {
			sqlWrite(r)
			code = "Kenteken verstuurd"
		}
	}
	data := PageData{Statuscode: code}
	tmpl.Execute(w, data)
}

func main() {
	fs := http.FileServer(http.Dir("./views")) // fileserver for css
	http.Handle("/", fs)

	http.HandleFunc("/registreer", index) // handle func index
	http.ListenAndServe(":80", nil)
}

func sqlWrite(r *http.Request) {

	db, err := sql.Open("mysql", "Jasper:p23hzbdjdsYm@tcp(capitaselectadb.mysql.database.azure.com:3306)/csdb?tls=true")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("INSERT INTO `csdb`.`kenteken` (`KentekenID`) VALUES ('" + r.FormValue("kenteken") + "')")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
}

// ga verder met json opslag https://dev.to/evilcel3ri/append-data-to-json-in-go-5gbj
// en html info uit textbox/form https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html
//https://blog.logrocket.com/dockerizing-go-application/
