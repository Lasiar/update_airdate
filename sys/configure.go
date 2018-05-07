package sys

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type config struct {
	ConnStr string `json:"connect_string"`
	Log     *log.Logger
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
	c.log = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
}
