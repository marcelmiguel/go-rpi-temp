package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	configuration := Configuration{}
	configuration.DefaultPath()
	configuration.Open("")

	if configuration.StartServices {
		start_services()
	}

	//var day [24]Temperature

	t := make(chan Temperature, 1)
	towm := make(chan Temperature, 1)
	dayc := make(chan []Temperature, 1)

	go readData(&configuration, t, dayc)
	go readDataOWM(&configuration, towm)

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, t)
	})*/

	fs := http.FileServer(http.Dir(configuration.ConfPath+"."))
	http.Handle("/", http.StripPrefix("/", fs))

	http.HandleFunc("/api/v1/temp", func(w http.ResponseWriter, r *http.Request) {
		handlerREST(w, r, t)
	})

	http.HandleFunc("/api/v1/tempowm", func(w http.ResponseWriter, r *http.Request) {
		handlerRESTOWM(w, r, towm)
	})

	http.HandleFunc("/api/v1/tempday", func(w http.ResponseWriter, r *http.Request) {
		handlerRESTDAY(w, r, dayc)
	})
	http.HandleFunc("/api/v1/os", func(w http.ResponseWriter, r *http.Request) {
		handlerRESTOS(w, r, &configuration)
	})

	port := strconv.Itoa(configuration.ServerPort)
	fmt.Println("web server on localhost:" + port)

	log.Fatal(http.ListenAndServe(":" + port, nil))

}
