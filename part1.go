package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func open(location string) {
	dat, err := ioutil.ReadFile(location)
	checkErr(err)
	fmt.Println(string(dat))
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	open(args[0])
}
