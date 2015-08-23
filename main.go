package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	configuration := Configuration{}
	configuration.Open("")

	if configuration.StartServices {
		start_services()
	}

	t := make(chan Temperature, 1)

	go readData(&configuration, t)

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, t)
	})*/

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", http.StripPrefix("/", fs))

	http.HandleFunc("/api/v1/temp", func(w http.ResponseWriter, r *http.Request) {
		handlerREST(w, r, t)
	})

	port := strconv.Itoa(configuration.ServerPort)
	fmt.Println("web server on localhost:" + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
