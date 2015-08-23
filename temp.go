package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Temperature struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
	Id         string `json:"id"`
	Valid      bool `json:"valid"`
	Modified   bool `json:"modified"`
	Location   string `json:"location"`
}

var temps struct {
	actual   Temperature
	previous Temperature
}

func ReadSensorFile(file string) Temperature {
	//fmt.Println("Filename ",file)
	base := filepath.Base(filepath.Dir(file))
	tmp := Temperature{}
	tmp.Valid = false
	tmp.Id = base[3:]

	inFile, err := os.Open(file)
	if err != nil {
		return tmp
	}

	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		linea := lines[i]

		//fmt.Println(linea)

		if i == 0 { //linea de estado
			//fmt.Println("Status ",linea[len(linea)-3:])
			if linea[len(linea)-3:] == "YES" {
				tmp.Valid = true
			}
		} else if i == 1 { //linea de temperatura
			li := strings.LastIndexAny(linea, "t=")
			if li >= 0 {
				celsius, err1 := strconv.ParseFloat(linea[li+1:], 64)
				if err1 != nil {
					tmp.Valid = false
				} else {
					tmp.Celsius = celsius / 1000.0
					tmp.Fahrenheit = tmp.Celsius*9.0/5.0 + 32.0 //TODO pasar a función del strcut
					tmp.Location = "Barcelona"
				}
			}
		}
	}

	return tmp
}

func readData(configuration *Configuration, t chan Temperature) {
	files1, _ := filepath.Glob(configuration.Pathw1 + "28*/w1_slave")

	for {
		if len(t) > 0 {
			<-t
		}
		for _, f := range files1 {
			tmp := ReadSensorFile(f)

			t <- tmp

			if temps.previous.Celsius != tmp.Celsius {
				fmt.Printf("%2.3fºC ", tmp.Celsius)
				tmp.Modified = true
			} else {
				fmt.Printf(".")
				tmp.Modified = false
			}
			temps.previous = tmp
		}

		time.Sleep(time.Duration(configuration.Frequency) * time.Second)
	}
}
