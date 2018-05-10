package sys

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type config struct {
	ConnStr string `json:"connect_string"`
	DB      string `json:"db"`
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

	if err := dc.Decode(&c); err != nil {
		log.Fatal("Read config file: ", err)
	}

	if c.DB == "" || c.ConnStr == "" {
		log.Fatal("Config file corrupted: ", c.ConnStr, c.DB)
	}

	c.Err = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	c.Warn = log.New(os.Stderr, "[WARNING] ", log.Ldate|log.Ltime)
}
