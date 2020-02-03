package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var debug = flag.Bool("debug", false, "Usage: -degbug=<val>")
var inputFile = flag.String("inputFile", "./input", "Usage: -inputFile=<path-to-input>")

func getVal(slice *[]int, index int) int {
	if index >= len(*slice) {
		tmp := make([]int, index-len(*slice)+1)
		*slice = append(*slice, tmp...)
	}
	return (*slice)[index]
}

func assignVal(slice *[]int, index int, val int) {
	if index >= len(*slice) {
		tmp := make([]int, index-len(*slice)+1)
		*slice = append(*slice, tmp...)
	}
	(*slice)[index] = val
}

func modeResolver(contents *[]int, pos int, mode string, relBase int) int {
	if mode == "0" {
		return getVal(contents, getVal(contents, pos))
	} else if mode == "1" {
		return getVal(contents, pos)
	} else {
		// mode == "2"
		return getVal(contents, relBase+getVal(contents, pos))
	}
}

func intcodeProcess(processID int, procPtr int, ampState []int, inSigVals []int, inSigPtr int, db int) (bool, int, []int, int, int) {

	var ampOut int
	exit, ret := false, false
	contents := make([]int, 0)
	contents = append(contents, ampState...)
	relBase := 0

	for i := procPtr; i < len(contents); {

		iUpdate := true
		op := strconv.Itoa(getVal(&contents, i))
		if db != 0 {
			fmt.Println("->", getVal(&contents, i))
		}
		op = strings.Repeat("0", 5-len(op)) + op

		opcode := op[3:]
		// a = c (op) b
		op_c := op[2:3]
		op_b := op[1:2]
		op_a := op[0:1]

		jump := 1
		switch opcode {
		// 0 -> position mode
		// 1 -> immediate mode
		// 2 -> relative mode
		case "01":
			// add
			c := modeResolver(&contents, i+1, op_c, relBase) + modeResolver(&contents, i+2, op_b, relBase)
			id := getVal(&contents, i+3)
			if op_a == "2" {
				id += relBase
			}
			assignVal(&contents, id, c)
			jump = 4

		case "02":
			//multiply
			c := modeResolver(&contents, i+1, op_c, relBase) * modeResolver(&contents, i+2, op_b, relBase)
			id := getVal(&contents, i+3)
			if op_a == "2" {
				id += relBase
			}
			assignVal(&contents, id, c)
			jump = 4

		case "03":
			//read
			if inSigPtr < len(inSigVals) {
				id := getVal(&contents, i+1)
				if op_c == "2" {
					id += relBase
				}
				// id := getVal(&contents, i+1)
				fmt.Println("ID IS", id, relBase)
				assignVal(&contents, id, inSigVals[inSigPtr])
				inSigPtr++
			} else {
				if db == 1 {
					fmt.Println("Halting for input signal. procId: ", processID, " | procPtr: ", procPtr)
				}
				ret = true
			}
			jump = 2

		case "04":
			//write
			c := modeResolver(&contents, i+1, op_c, relBase)
			fmt.Println("i:", i, " | ampOutput:", c)
			ampOut = c
			// ret = true
			jump = 2

		case "05":
			//jump-if-true
			c := modeResolver(&contents, i+1, op_c, relBase)
			if c != 0 {
				// if op_b == "2"
				i = modeResolver(&contents, i+2, op_b, relBase)
				iUpdate = false
			}
			jump = 3

		case "06":
			//jump-if-false
			c := modeResolver(&contents, i+1, op_c, relBase)
			if c == 0 {
				i = modeResolver(&contents, i+2, op_b, relBase)
				iUpdate = false
			}
			jump = 3

		case "07":
			//less-than
			tpFlag := modeResolver(&contents, i+1, op_c, relBase) < modeResolver(&contents, i+2, op_b, relBase)
			id := getVal(&contents, i+3)
			if op_a == "2" {
				id += relBase
			}
			if tpFlag {
				assignVal(&contents, id, 1)
			} else {
				assignVal(&contents, id, 0)
			}
			jump = 4

		case "08":
			//equals
			tpFlag := modeResolver(&contents, i+1, op_c, relBase) == modeResolver(&contents, i+2, op_b, relBase)
			id := getVal(&contents, i+3)
			if op_a == "2" {
				id += relBase
			}
			if tpFlag {
				assignVal(&contents, id, 1)
			} else {
				assignVal(&contents, id, 0)
			}
			jump = 4

		case "09":
			//update relative base
			relBase = relBase + modeResolver(&contents, i+1, op_c, relBase)
			jump = 2

		case "99":
			if db != 0 {
				fmt.Println("Terminating. Exit code 1.")
			}
			exit = true

		default:
			fmt.Println("Encountered invalid opcode. Val: ", op, contents[i])
		}

		if iUpdate {
			i += jump
		}
		if exit == true || ret == true {
			break
		}
	}
	return exit, procPtr, contents, ampOut, inSigPtr
}

func main() {
	flag.Parse()
	f, _ := ioutil.ReadFile(*inputFile)
	ip := strings.Split(string(f), ",")

	var contents []int
	for i := 0; i < len(ip); i++ {
		val, _ := strconv.Atoi(ip[i])
		contents = append(contents, val)
	}

	if *debug == true {
		fmt.Println("Len:", len(contents), "\nContents:", contents)
	}
	// fmt.Println(math.MaxInt64)
	inSigVals := []int{2}
	intcodeProcess(0, 0, contents, inSigVals, 0, 1)
	// fmt.Println(sth)
}
