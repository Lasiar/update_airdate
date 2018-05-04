package sys

import (
	"os"
	"log"
	"encoding/json"
	"sync"
)

type config struct {
	ConnStr string `json:"connect_string"`
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

func (c *config)load() {
	confFile, err := os.Open("config.json")
	if err != nil{
		log.Fatal(err)
	}
	dc := json.NewDecoder(confFile)
	dc.Decode(&c)
}