package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func open(location string) {
	data, err := ioutil.ReadFile(location)
	checkErr(err)
	return data
}

func parse(data byte[]) {
	
	for item := range data {
		fmt.PrintLn(string(item))
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	data = open(args[0])
	outputName = args[1]

}
