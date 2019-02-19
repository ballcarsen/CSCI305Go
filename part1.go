package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func open(location string) []byte {
	data, err := ioutil.ReadFile(location)
	checkErr(err)
	return data
}

func parse(data []byte) {
	fmt.Println(string(data))
	for item := range data {
		fmt.Println(string(item))
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	data := open(args[0])
	outputName := args[1]
	fmt.Println(outputName)
	parse(data)
}
