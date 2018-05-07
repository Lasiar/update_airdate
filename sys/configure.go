package sys

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type config struct {
	ConnStr string `json:"connect_string"`
	ErrDB   *log.Logger
	ErrWeb  *log.Logger
	WarnWeb *log.Logger
}

var _config *config
var _once sync.Once

func GetConfig() *config {
	_once.Do(func() {
		_config = new(config)
		_config.load()
	})
	return _config
}

func (c *config) load() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	confFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	dc := json.NewDecoder(confFile)
	dc.Decode(&c)
	c.ErrDB = log.New(os.Stderr, "[ERROR] database ", log.Ldate|log.Ltime|log.Lshortfile)
	c.ErrWeb = log.New(os.Stderr, "[ERROR] web ", log.Ldate|log.Ltime|log.Lshortfile)
	c.WarnWeb = log.New(os.Stderr, "[WARNING] web ", log.Ldate|log.Ltime)
}
