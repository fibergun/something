package main

import (
	"log"
	"net/http"
	"something/pages"
)

func main() {

	http.HandleFunc("/save/", pages.Save)
	http.HandleFunc("/edit/", pages.Edit)
	http.HandleFunc("/view/", pages.View)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
