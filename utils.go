package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

func execute_cmd(cmd string, wg *sync.WaitGroup) {
	//fmt.Println("command is ", cmd)
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	_, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Print("Error ", err)
	}
	//fmt.Println(out)
	wg.Done()
}
