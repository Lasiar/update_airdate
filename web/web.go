package web

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
	"log"
	"io"
	"time"
)

type updateRequest struct {
	DateOneStart  string `json:"date_one_start"`
	DateOneFinish string `json:"date_one_finish"`
}

type responseRequest struct {
	Success bool `json:"success"`
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("запрет")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var upReq updateRequest

	err = json.Unmarshal([]byte(buf), &upReq)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	time.Sleep(20*time.Second)
	fmt.Println(upReq)
	fmt.Fprint(w, upReq)
}

func HandleFront(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("assets/index.html")
	if err != nil {
		log.Println(err)
	}
	io.Copy(w, f)
}

