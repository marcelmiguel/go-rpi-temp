package conf

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
}

func Open() Configuration {
	configuration := Configuration{}

	file, err := os.Open("conf.json")
	if err != nil {
		return configuration
	}
	decoder := json.NewDecoder(file)
	err1 := decoder.Decode(&configuration)
	if err1 != nil {
		fmt.Println("error reading conf file:", err)
	}

	if configuration.Frequency == 0 {
		configuration.Frequency = 3
	}

	return configuration
}
