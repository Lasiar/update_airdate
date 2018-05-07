package main

import (
	"net/http"
	"kre_air_update/web"
	"log"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("assets/css/"))))
	http.Handle("/fig/", http.StripPrefix("/fig/", http.FileServer(http.Dir("assets/fig/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("assets/js/"))))
	http.HandleFunc("/update", web.HandleUpdate)
	http.HandleFunc("/", web.HandleFront)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
