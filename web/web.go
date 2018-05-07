package web

import (
	"encoding/json"
	"fmt"
	"kre_air_update/model"
	"kre_air_update/sys"
	"net/http"
	"time"
)

type request struct {
	Who        string `json:"who"`
	DateStart  string `json:"date_start"`
	DateFinish string `json:"date_finish"`
}

const (
	startRow = 1
	midleRow = 16
	endRow   = 18
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
		fmt.Println(err)
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

	var req request

	dc := json.NewDecoder(r.Body)
	err := dc.Decode(&req)
	if err != nil {
		sys.GetConfig().Log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(req.DateFinish)+len(req.DateStart) < 2 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.Who != "pb" && req.Who != "all" {
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
		if tmStart.Time != tmFinish.Time {
			if tmStart.After(tmFinish.Time) {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}

	en := json.NewEncoder(w)
	if err := req.Update(tmStart.Time, tmFinish.Time); err != nil {
		sys.GetConfig().Log.Println(err)
		en.Encode(responseRequest{Success: false})
	} else {
		en.Encode(responseRequest{Success: true})
	}
}

func HandleFront(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/index.html")
}
