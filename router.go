package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func handlerREST(w http.ResponseWriter, r *http.Request, t <-chan Temperature) {
	if len(t) > 0 {
		temps.actual = <-t
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(temps.actual); err != nil {
		fmt.Fprintf(w, "Error", err)
	}
}

func handlerRESTOWM(w http.ResponseWriter, r *http.Request, t <-chan Temperature) {
	if len(t) > 0 {
		temps.owm = <-t
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(temps.owm); err != nil {
		fmt.Fprintf(w, "Error", err)
	}
}

func handlerRESTDAY(w http.ResponseWriter, r *http.Request, dayc <-chan []Temperature) {
	if len(dayc) > 0 {
		//day := <-dayc
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	/*if err := json.NewEncoder(w).Encode(temps.actual); err != nil {
		fmt.Fprintf(w, "Error", err)
	}*/

	json := "[ [\"Hour\", \"ºC\"], [\"09h\", 24.56],[\"10h\", 26.34],[\"11h\", 27.12],[\"12h\", 27.32],[\"13h\", 27.52]]"

	fmt.Fprintf(w, json)
}

func handlerRESTOS(w http.ResponseWriter, r *http.Request, conf *Configuration) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, conf.OS)
}