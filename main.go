package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	files, _ := ioutil.ReadDir("./")
	for _, f:= range files {
		fmt.Println(f.Name())
	}
	
	files1, _ := filepath.Glob("*")
	fmt.Println(files1)
}