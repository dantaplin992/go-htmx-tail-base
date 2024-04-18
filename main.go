package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	fmt.Println("Hello World!")

	var port string

	if os.Getenv("PORT") == "" {
		port = "8000"
	} else {
		port = os.Getenv("PORT")
	}
	portWithPrefix := fmt.Sprintf(":%s", port)

	fmt.Println("using port: ", portWithPrefix)

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/get_clicked", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("partials/clicked.html"))

		tmp.Execute(w, nil)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
