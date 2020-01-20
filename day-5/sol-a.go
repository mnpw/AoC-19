package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "./input", "Input for solver")
var debug = flag.Bool("debug", false, "Set true for debug output")

func modeResolver(contents []int, pos int, mode string) int {
	if mode == "1" {
		return contents[pos]
	}
	return contents[contents[pos]]
}

func intcodeProcess(contents []int) int {
	exit := false

	for i := 0; i < len(contents); {

		iUpdate := true
		op := strconv.Itoa(contents[i])
		fmt.Println("->", contents[i])
		op = strings.Repeat("0", 5-len(op)) + op

		opcode := op[3:]
		// a = c (op) b
		op_c := op[2:3]
		op_b := op[1:2]

		jump := 1
		switch opcode {
		// 0 -> position mode
		// 1 -> immediate mode
		case "01":
			// add
			c := modeResolver(contents, i+1, op_c) + modeResolver(contents, i+2, op_b)
			contents[contents[i+3]] = c
			jump = 4

		case "02":
			//multiply
			c := modeResolver(contents, i+1, op_c) * modeResolver(contents, i+2, op_b)
			contents[contents[i+3]] = c
			jump = 4

		case "03":
			//read
			contents[contents[i+1]] = 5
			jump = 2

		case "04":
			//write
			fmt.Println(contents[contents[i+1]])
			jump = 2

		case "05":
			//jump-if-true
			c := modeResolver(contents, i+1, op_c)
			if c != 0 {
				i = modeResolver(contents, i+2, op_b)
				iUpdate = false
			}
			jump = 3

		case "06":
			//jump-if-false
			c := modeResolver(contents, i+1, op_c)
			if c == 0 {
				i = modeResolver(contents, i+2, op_b)
				iUpdate = false
			}
			jump = 3

		case "07":
			//less-than
			tpFlag := modeResolver(contents, i+1, op_c) < modeResolver(contents, i+2, op_b)
			if tpFlag {
				contents[contents[i+3]] = 1
			} else {
				contents[contents[i+3]] = 0
			}
			jump = 4

		case "08":
			//equals
			tpFlag := modeResolver(contents, i+1, op_c) == modeResolver(contents, i+2, op_b)
			if tpFlag {
				contents[contents[i+3]] = 1
			} else {
				contents[contents[i+3]] = 0
			}
			jump = 4

		case "99":
			fmt.Println("Terminating. Exit code 1.")
			exit = true

		default:
			fmt.Println("Encountered invalid opcode. Val: ", op, contents[i])
		}

		if iUpdate {
			i += jump
		}
		if exit == true {
			break
		}
	}
	return 1
}

func main() {

	flag.Parse()
	f, _ := ioutil.ReadFile(*inputFile)
	_contents := strings.Split(string(f), ",")

	var contents []int
	for i := 0; i < len(_contents); i++ {
		val, _ := strconv.Atoi(_contents[i])
		contents = append(contents, val)
	}

	if *debug {
		fmt.Println("Type of contents is ", reflect.TypeOf(contents))
		fmt.Println("The len of contents if", len(contents))
		fmt.Println("contents: \n", contents)
	}
	intcodeProcess(contents)

}
