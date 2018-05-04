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

type request struct {
	Who        string `json:"who"`
	DateStart  string `json:"date_one_start"`
	DateFinish string `json:"date_one_finish"`
}

const (
	startRow = 1
	midleRow = 16
	endRow = 18
)


func (u request) Update(start, finish time.Time) error {
	if u.Who == "pb" {
		return model.Update(start, finish, midleRow, endRow)
	} else {
		return model.Update(start, finish, startRow, midleRow)
	}
	return nil
}

type responseRequest struct {
	Success bool `json:"success"`
}

type date struct {
	time.Time
}

func (d *date) Parse(str string) bool {
	tmp, err := time.Parse("2006-01-02", str)
	if err != nil {
		return false
	}
	d.Time = tmp
	return true
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var req request

	err = json.Unmarshal([]byte(buf), &req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var tmStart, tmFinish date

	if !tmStart.Parse(req.DateStart) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.DateFinish == "" {
		tmFinish = tmStart
	} else {
		if !tmFinish.Parse(req.DateFinish) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if req.Who != "pb" && req.Who != "all" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	en := json.NewEncoder(w)
	if err := req.Update(tmStart.Time, tmFinish.Time); err != nil {
		fmt.Println(err)
		en.Encode(responseRequest{Success: false})
	} else {
		en.Encode(responseRequest{Success: true})
	}
}

func HandleFront(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("assets/index.html")
	if err != nil {
		log.Println(err)
	}
	io.Copy(w, f)
}
