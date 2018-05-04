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
	"air/model"
)

type updateRequest struct {
	Who           string `json:"who"`
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
		return
	}

	fmt.Println(upReq.DateOneStart)

	tm_start, err := time.Parse("2006-01-02", upReq.DateOneStart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var tm_finish time.Time

	if upReq.DateOneFinish == "" {
		tm_finish = tm_start
	} else {
		if tm_finish, err = time.Parse("2006-01-02", upReq.DateOneFinish); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if upReq.Who != "pb" && upReq.Who != "all" {
		w.WriteHeader(http.StatusInternalServerError)
	}




	if upReq.Who == "pb" {
		en := json.NewEncoder(w)
		if err := model.UpdatePb(tm_start, tm_finish); err != nil {
			fmt.Println(err)
			en.Encode(responseRequest{Success: false})
		} else {
			en.Encode(responseRequest{Success: true})
		}
	}

	if upReq.Who == "all" {
		en := json.NewEncoder(w)
		if err := model.UpdateAll(tm_start, tm_finish); err != nil {
			fmt.Println(err)
			en.Encode(responseRequest{Success: false})
		} else {
			en.Encode(responseRequest{Success: true})
		}
	}
}

func HandleFront(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("assets/index.html")
	if err != nil {
		log.Println(err)
	}
	io.Copy(w, f)
}
