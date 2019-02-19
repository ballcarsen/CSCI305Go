package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func open(location string) []byte {
	data, err := ioutil.ReadFile(location)
	checkErr(err)
	return data
}

func write(outLoc string, data []byte) {
	err := ioutil.WriteFile(outLoc, data, 0644)
	checkErr(err)
}

func parse(data []byte) []byte {
	var output []byte
	for i := range data {
		chr := string(data[i])
		if _, err := strconv.Atoi(chr); err != nil {
			output = append(output, data[i])
		}
	}
	return output
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
	dt := parse(data)
	write(outputName, dt)
}
