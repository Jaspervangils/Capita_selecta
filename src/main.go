package main

import (
	"html/template"
	"net/http"
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

// ga verder met json opslag https://dev.to/evilcel3ri/append-data-to-json-in-go-5gbj
// en html info uit textbox/form https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html

//https://blog.logrocket.com/dockerizing-go-application/
