package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	PlotlyAPIToken string
	PlotlySecret   string
	Frequency      int
	Pathw1         string
	ServerPort     int
	StartServices  bool
	Log            bool
}

func (conf *Configuration) Open(FileName string) bool {

	if FileName=="" {
		FileName = "conf.json"
	}

	file, err := os.Open(FileName)
	if err != nil {
		return false
	}

	conf.StartServices = false
	conf.ServerPort = 80
	conf.Frequency = 3

	decoder := json.NewDecoder(file)
	err1 := decoder.Decode(&conf)
	if err1 != nil {
		if conf.Log {
			fmt.Println("error reading conf file:", err)
		}
	}

	if conf.Frequency == 0 {
		conf.Frequency = 1
	}

	if conf.Log {
		fmt.Println(conf)
	}

	return true
}

