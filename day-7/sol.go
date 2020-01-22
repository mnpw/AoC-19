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

func modeResolver(contents []int, pos int, mode string) int {
	if mode == "1" {
		return contents[pos]
	}
	return contents[contents[pos]]
}

func intcodeProcess(contents []int, phase int, input int, db int, insPtr ...int) int {
	var ampOut int
	exit := false
	applyPhase := true
	//applyInput is !applyPhase
	for i := 0; i < len(contents); {

		iUpdate := true
		op := strconv.Itoa(contents[i])
		if db != 0 {
			fmt.Println("->", contents[i])
		}
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
			if applyPhase {
				contents[contents[i+1]] = phase
				applyPhase = false
			} else {
				contents[contents[i+1]] = input
			}
			jump = 2

		case "04":
			//write
			if db != 0 {
				fmt.Println("ampOutput is ", contents[contents[i+1]])
			}
			ampOut = contents[contents[i+1]]
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
		if exit == true {
			// insPtr = i
			break
		}
	}
	return ampOut
}

func modSlice(sl *[]string) {
	(*sl)[0] = "99"
	*sl = append(*sl, "69")
	fmt.Println(len(*sl), *sl)
}

func genPerm(elements []string) []string {
	// someStringSlice := make([]string, 4)
	var someStringSlice = []string{"0"}
	fmt.Println(someStringSlice)
	// someStringSlice[0:] = []string{"0", "1", "2"}
	fmt.Println(someStringSlice)
	modSlice(&someStringSlice)
	fmt.Println(someStringSlice)
	return make([]string, 0)
}

func removeIndex(sl []string, id int) []string {
	elementsCopy := make([]string, 0)
	elementsCopy = append(elementsCopy, sl[:id]...)
	return append(elementsCopy, sl[id+1:]...)
}

func genPermutationRoutine(size int, currPermuation []string, elements []string, output *[][]string) {
	if len(currPermuation) == size {
		*output = append((*output), currPermuation)
		return
	}
	for i := 0; i < len(elements); i++ {
		tmpCurrPermutation := append(currPermuation, elements[i])
		tmpElements := removeIndex(elements, i)
		genPermutationRoutine(size, tmpCurrPermutation, tmpElements, output)
	}
	return
}

func genPermutation(elements []string) [][]string {
	output := make([][]string, 0)
	genPermutationRoutine(len(elements), make([]string, 0), elements, &output)
	return output
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

	var tmp = []string{"0", "1", "2", "3", "4"}
	permutations := genPermutation(tmp)
	// fmt.Println(len(permutations), permutations)
	maxOutput := 0
	for i := 0; i < len(permutations); i++ {
		currInput := 0
		for j := 0; j < 5; j++ {
			phase, _ := strconv.Atoi(permutations[i][j])
			currInput = intcodeProcess(contents, phase, currInput, 0)
		}
		if maxOutput < currInput {
			maxOutput = currInput
		}
	}
	fmt.Println(maxOutput)
}
