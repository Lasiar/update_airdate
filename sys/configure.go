package sys

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type config struct {
	ConnStr string `json:"connect_string"`
	Err     *log.Logger
	Warn    *log.Logger
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
	c.Err = log.New(os.Stderr, "[ERROR  ] ", log.Ldate|log.Ltime|log.Lshortfile)
	c.Warn = log.New(os.Stderr, "[WARNING ] ", log.Ldate|log.Ltime)
}
