package main

import (
	"fmt"
	"log"
	"something/pages"
)

func main() {
	p1 := &pages.Page{Title: "Page1", Body: []byte("Hello World")}
	p1.Save()
	p2, err := pages.LoadPage("Page1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p2)
}
