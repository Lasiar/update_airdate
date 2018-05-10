package web

import (
	"encoding/json"
	"fmt"
	"kre_air_update/model"
	"kre_air_update/sys"
	"net/http"
	"time"
)

const (
	startRow = 1
	midleRow = 16
	endRow   = 18
)

type request struct {
	Who        string `json:"who"`
	DateStart  string `json:"date_start"`
	DateFinish string `json:"date_finish"`
}

func (u request) Update(start, finish time.Time) error {
	if u.Who == "pb" {
		return model.Update(start, finish, midleRow, endRow)
	} else {
		return model.Update(start, finish, startRow, midleRow)
	}
	return nil
}

func (u request) String() string {
	return fmt.Sprintf("%#v", u)
}

type myStr string

func (u myStr) String() string {
	return string(u)
}

func (u myStr) Error() string {
	return fmt.Sprintf("%s", string(u))
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

type responseRequest struct {
	Success bool `json:"success"`
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		printErr(r, myStr("wrong method"), myStr(""))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req request
	dc := json.NewDecoder(r.Body)
	err := dc.Decode(&req)
	if err != nil {
		printErr(r, err, myStr("decode json"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(req.DateFinish)+len(req.DateStart) < 2 {
		printErr(r, myStr("wrong len date"), req)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if req.Who != "pb" && req.Who != "all" {
		printErr(r, myStr("wrong method"), req)
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
				printErr(r, myStr("wrong date"), req)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}

	en := json.NewEncoder(w)

	if err := req.Update(tmStart.Time, tmFinish.Time); err != nil {
		en.Encode(responseRequest{Success: false})
		sys.GetConfig().Err.Println(err)
		return
	}
	en.Encode(responseRequest{Success: true})
}

func HandleFront(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/index.html")
}

func printErr(r *http.Request, err error, req fmt.Stringer) {
	sys.GetConfig().Warn.Printf("FROM %v %v;  DATA: %v", r.RemoteAddr, err, req)
}
