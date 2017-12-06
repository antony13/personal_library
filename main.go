package main

import (
	"fmt"
	"net/http"
	"html/template" //add the template library
)

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
		fmt.Fprint(w, "Hello world!!!!")
	})

	fmt.Println(http.ListenAndServe(":8000",nil))
}

