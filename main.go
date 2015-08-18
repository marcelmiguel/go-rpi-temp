package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os/exec"
	"sync"
	"strings"
)

func exe_cmd(cmd string, wg *sync.WaitGroup) {
  fmt.Println("command is ",cmd)
  // splitting head => g++ parts => rest of the command
  parts := strings.Fields(cmd)
  head := parts[0]
  parts = parts[1:len(parts)]

  out, err := exec.Command(head,parts...).Output()
  if err != nil {
    fmt.Printf("%s", err)
  }
  fmt.Printf("%s", out)
  wg.Done() // Need to signal to waitgroup that this goroutine is done
}

func main() {
	files, _ := ioutil.ReadDir("./")
	for _, f:= range files {
		fmt.Println(f.Name())
	}
	
	files1, _ := filepath.Glob("*")
	fmt.Println(files1)
	
	//testing execution of exe
	wg := new(sync.WaitGroup)
    wg.Add(1)
    go exe_cmd("curl", wg)
    wg.Wait()
	
}