package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)


/*func handler(w http.ResponseWriter, r *http.Request, t <-chan Temperature) {
	if len(t) > 0 {
		temps.actual = <-t
	}
	if temps.actual.Valid {
		fmt.Fprintf(w, "Temperature %s : %gºc", temps.actual.Id, temps.actual.Celsius)
	}
}*/

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

