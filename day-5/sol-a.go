package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "./input", "Input for solver")
var debug = flag.Bool("debug", false, "Set true for debug output")

func main() {
	flag.Parse()
	f, _ := ioutil.ReadFile(*inputFile)
	contents := strings.Split(string(f), ",")

	if *debug {
		fmt.Println("The len of contents if", len(contents), contents[0])
		for index := 0; index < len(contents); index++ {
			fmt.Print(contents[index])
		}
	}
}
