package main

import (
	"sync"
)

func start_services() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go execute_cmd("modprobe w1-gpio", wg)
	go execute_cmd("modprobe w1-therm", wg)
	wg.Wait()
}

