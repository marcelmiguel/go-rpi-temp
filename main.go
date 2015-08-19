package main

import (
	"fmt"
	"bufio"
	"path/filepath"
	"os"
	"os/exec"
	"sync"
	"strings"
	"strconv"
	"go-rpi-temp/utils"
	"net/http"
	"time"
)

type Temperature struct {
    Celsius   float64
    Farenheit float64
	Id string //serial number
	Valid bool
}

var temp Temperature

func exe_cmd(cmd string, wg *sync.WaitGroup) {
  fmt.Println("command is ",cmd)
  // splitting head => g++ parts => rest of the command
  parts := strings.Fields(cmd)
  head := parts[0]
  parts = parts[1:len(parts)]

  out, err := exec.Command(head,parts...).Output()
  if err != nil {
    fmt.Print(err)
  }
  fmt.Println(out)
  wg.Done() // Need to signal to waitgroup that this goroutine is done
}

func ReadTemp(file string) Temperature {
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

  for i:=0; i<len(lines); i++{
	linea := lines[i]
	
	//fmt.Println(linea)
	
	if i==0 { //linea de estado
	  //fmt.Println("Status ",linea[len(linea)-3:])
	  if linea[len(linea)-3:] == "YES" {
		tmp.Valid = true
	  }
	} else if i==1 { //linea de temperatura
	  li := strings.LastIndexAny(linea,"t=")
	  if li>=0 {
	    celsius, err1 := strconv.ParseFloat(linea[li+1:], 64)
		if err1 != nil {
			tmp.Valid = false
		} else {
			tmp.Celsius = celsius / 1000.0
			tmp.Farenheit = tmp.Celsius * 9.0 / 5.0 + 32.0 //TODO pasar a función del strcut			
		}		
	  }
	}
  }

  return tmp
}

func handler(w http.ResponseWriter, r *http.Request, t <-chan Temperature) {
	if len(t)>0 {
		temp = <- t
	}
	if temp.Valid {
		fmt.Fprintf(w, "Temperature %s : %g", temp.Id, temp.Celsius)	
	}
}

func readData(configuration *conf.Configuration, t chan Temperature) {
	files1, _ := filepath.Glob(configuration.Pathw1+"28*/w1_slave")

    for {
		if len(t)>0 {
			<- t
		}
		for _, f:= range files1 {
			tmp := ReadTemp(f)
			
			t <- tmp
			fmt.Println(tmp)
		}
		
		time.Sleep(time.Duration(configuration.Frequency) * time.Second)
    }
}
	
func main() {
	//execute modprobes
	wg := new(sync.WaitGroup)
    wg.Add(2)
    go exe_cmd("modprobe w1-gpio", wg)
	go exe_cmd("modprobe w1-therm", wg)
    wg.Wait()

	configuration := conf.Open()
	fmt.Println(configuration)

	t := make(chan Temperature, 1)

	go readData(&configuration, t)

	fmt.Println("web server on localhost:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
              handler(w, r, t)
       })
    http.ListenAndServe(":8080", nil)
	
}
