package main

import (
	"fmt"
	"net/http"
	"html/template" //add the template library
)

type Page struct {
	Name string 
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		par := Page{Name: "Antony"}
		if name := r.FormValue("name"); name!="" {
			par.Name = name
		}//end if

		if err := templates.ExecuteTemplate(w, "index.html", par); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
		fmt.Fprint(w, "Hello world!!!!")
	})

	fmt.Println(http.ListenAndServe(":8000",nil))
}

