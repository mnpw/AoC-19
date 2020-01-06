package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

var inputFile = flag.String("inputFile", "./input", "Input for solver")
var debug = flag.Bool("debug", false, "Set true for debug output")

func intcodeProcess(contents []string) int {
	for i, val := range contents {
		// opcode := val[len(val)-1]
		fmt.Print(reflect.TypeOf(val), ", ", i, ", ", val, " ", len(val), " ", []byte(val), "\n")
	}
	return 1
}

func main() {
	flag.Parse()
	f, _ := ioutil.ReadFile(*inputFile)
	contents := strings.Split(string(f), ",")

	if *debug {
		fmt.Print("Type of contents is ", reflect.TypeOf(contents))
		fmt.Println("The len of contents if", len(contents), contents[0])
		for index := 0; index < len(contents); index++ {
			fmt.Print(contents[index], ", ")
		}
	}
	intcodeProcess(contents)
}
