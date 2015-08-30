package main

import (
	"encoding/json"
	"fmt"
	"os"
	"bufio"
	"runtime"
)

type Configuration struct {
	PlotlyAPIToken string
	PlotlySecret   string
	Frequency      int
	Pathw1         string
	ServerPort     int
	StartServices  bool
	Log            bool
	ConfPath       string
	City           string
	LocalDescription string
	OS             string
}

func (conf *Configuration) Open(FileName string) bool {

	if FileName=="" {
		FileName = conf.ConfPath+"conf.json"
	}

	fmt.Println(FileName)
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

	conf.OS = runtime.GOOS+" "+runtime.GOARCH

	if conf.Log {
		fmt.Println(conf)
	}

	return true
}

func (conf *Configuration) DefaultPath() string {

	file, err := os.Open("/etc/weather.station")
	if err != nil {
		return ""
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		fmt.Println(lines)
	}

	if len(lines)>0 {
		conf.ConfPath = lines[0] + "/"
	} else {
		conf.ConfPath = ""
	}
	file.Close()
	return conf.ConfPath
}
