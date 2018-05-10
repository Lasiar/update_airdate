package main

import (
	"kre_air_update/sys"
	"kre_air_update/web"
	"net/http"
)

func main() {
	http.Handle("/fig/", http.StripPrefix("/fig/", http.FileServer(http.Dir("assets/fig/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("assets/js/"))))
	http.HandleFunc("/update", web.HandleUpdate)
	http.HandleFunc("/", web.HandleFront)
	sys.GetConfig().Err.Fatal(http.ListenAndServe(sys.GetConfig().Port, nil))
}
