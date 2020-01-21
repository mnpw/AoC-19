package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var debug = flag.Bool("debug", false, "--debug=<val>")
var inputFile = flag.String("inputFile", "./input", "-inputFile=<filename>")

func main() {
	flag.Parse()
	f, _ := ioutil.ReadFile(*inputFile)
	ip := strings.Split(string(f), ",")

	var contents []int
	for i := 0; i < len(ip); i++ {
		val, _ := strconv.Atoi(ip[i])
		contents = append(contents, val)
	}

	fmt.Println(contents)
}
