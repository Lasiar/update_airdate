package web

import (
	"encoding/json"
	"kre_air_update/model"
	"net/http"
	"time"
	"kre_air_update/sys"
	"fmt"
	"strings"
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

func (d *date) Parse(str string) (err error) {
	d.Time, err = time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	return err
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		printErr(r, "wrong method", "")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req request

	dc := json.NewDecoder(r.Body)
	err := dc.Decode(&req)
	if err != nil {
		printErr(r, err, "decode json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(req.DateFinish)+len(req.DateStart) < 2 {
		printErr(r, "wrong len date", req)
		sys.GetConfig().WarnWeb.Println("wrong len time")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.Who != "pb" && req.Who != "all" {
		printErr(r, "non existent type", req)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var tmStart, tmFinish date

	if err := tmStart.Parse(req.DateStart); err != nil {
		printErr(r, err, req)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.DateFinish == "" {
		tmFinish = tmStart
	} else {
		if err := tmFinish.Parse(req.DateFinish); err != nil {
			printErr(r, err, req)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if tmStart.Time != tmFinish.Time {
			if tmStart.After(tmFinish.Time) {
				printErr(r, "wrong date", req)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
	fmt.Println(tmStart, tmFinish)
	en := json.NewEncoder(w)
	if err := req.Update(tmStart.Time, tmFinish.Time); err != nil {
		en.Encode(responseRequest{Success: false})
		sys.GetConfig().ErrDB.Println(err)
	} else {
		en.Encode(responseRequest{Success: true})
	}
}

func HandleFront(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/index.html")
}

func printErr(r *http.Request, err interface{}, req interface{}) {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	sys.GetConfig().WarnWeb.Printf("FROM %v %v;  DATA: %#v", ip, err, req)
}