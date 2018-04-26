package main

import (
	"net/http"
	"air/web"
	"log"
)

func main() {
	http.Handle("/fig/", http.StripPrefix("/fig/", http.FileServer(http.Dir("fig/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	http.HandleFunc("/update", web.HandleUpdate)
	http.HandleFunc("/", web.HandleFront)
	log.Fatal(http.ListenAndServe(":8080", nil))
}