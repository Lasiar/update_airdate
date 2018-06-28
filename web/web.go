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

const (
	startDateAll   = iota
	finishDateAll
	startDatePb
	finishDatePb
)

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

func HandleGetDate(w http.ResponseWriter, r *http.Request) {


	if r.Method != http.MethodPost {
		printErr(r.RemoteAddr, myStr("wrong method"), myStr(r.Method))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	dt, err := model.Select()
	if err != nil {
		printErr(r.RemoteAddr, err, nil)
	}
	fmt.Println(dt)
	req := new([2]request)
	req[0] = request{"all", dt[startDateAll], dt[finishDateAll]}
	req[1] = request{"pb", dt[startDatePb], dt[finishDatePb]}
	dec := json.NewEncoder(w)
	dec.Encode(req)
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		printErr(r.RemoteAddr, myStr("wrong method"), myStr(r.Method))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req request
	dc := json.NewDecoder(r.Body)
	if err := dc.Decode(&req); err != nil {
		printErr(r.RemoteAddr, err, myStr("decode json"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(req.DateFinish)+len(req.DateStart) < 2 {
		printErr(r.RemoteAddr, myStr("wrong len date"), req)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if req.Who != "pb" && req.Who != "all" {
		printErr(r.RemoteAddr, myStr("wrong method"), req)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tmStart, tmFinish date

	if err := tmStart.Parse(req.DateStart); err != nil {
		printErr(r.RemoteAddr, err, req)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.DateFinish == "" || req.DateFinish == req.DateStart {
		tmFinish = tmStart
	} else {
		if err := tmFinish.Parse(req.DateFinish); err != nil {
			printErr(r.RemoteAddr, err, req)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if tmStart.After(tmFinish.Time) {
			printErr(r.RemoteAddr, myStr("wrong date"), req)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	en := json.NewEncoder(w)

	if err := req.Update(tmStart.Time, tmFinish.Time); err != nil {
		en.Encode(responseRequest{Success: false})
		sys.GetConfig().Err.Println(err)
		return
	}
	sys.GetConfig().Info.Printf("update DB [FROM] %v [DATA] %v", r.RemoteAddr, req)
	en.Encode(responseRequest{Success: true})
}

func HandleFront(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/index.html")
}

func printErr(ipAddr string, err error, req fmt.Stringer) {
	sys.GetConfig().Warn.Printf("%v [FROM] %v [DATA] %v", err, ipAddr, req)
}
