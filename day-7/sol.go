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

func intcodeProcess(id int, insPtr int, ampState []int, input [][]int, inputPtr int, db int) (bool, int, []int, int, int) {
	var ampOut int
	exit := false
	ret := false
	contents := make([]int, 0)
	contents = append(contents, ampState...)
	//applyInput is !usePhase
	for i := insPtr; i < len(contents); {
		// time.Sleep(1000 * time.Millisecond)
		// fmt.Println(contents)
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
			if inputPtr < len(input[id]) {
				contents[contents[i+1]] = input[id][inputPtr]
				inputPtr++
			} else {
				break
			}
			jump = 2

		case "04":
			//write
			if db != 0 {
				fmt.Println("ampOutput is ", contents[contents[i+1]])
			}
			ampOut = contents[contents[i+1]]
			ret = true
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
		insPtr = i
		if exit == true || ret == true {
			break
		}
	}
	return exit, insPtr, contents, ampOut, inputPtr
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
	fmt.Println(len(contents))
	ampCount := 5
	var tmp = []string{"5", "6", "7", "8", "9"}

	permutations := genPermutation(tmp)
	// permutations := make([][]string, 0)
	// permutations = append(permutations, []string{"9", "8", "7", "6", "5"})
	maxOutput := 0

	for i := 0; i < len(permutations); i++ {
		// fmt.Println("i:", i)
		output := 0

		ampStates := make([][]int, ampCount)
		ampPtrs := make([]int, ampCount)
		ampIps := make([][]int, ampCount)
		ampIpId := make([]int, ampCount)
		for j := 0; j < ampCount; j++ {
			ampStates[j] = make([]int, len(contents))
			copy(ampStates[j], contents)
			ampPtrs[j] = 0
			ampIps[j] = make([]int, 1)
			phase, _ := strconv.Atoi(permutations[i][j])
			ampIps[j][0] = phase
			if j == 0 {
				ampIps[0] = append(ampIps[0], 0)
			}
			ampIpId[j] = 0
		}

		x := false

		for j := 0; ; j++ {
			id := j % 5
			nextId := (j + 1) % 5
			// time.Sleep(1000 * time.Millisecond)
			x, ampPtrs[id], ampStates[id], output, ampIpId[id] = intcodeProcess(id, ampPtrs[id], ampStates[id], ampIps, ampIpId[id], 0)
			ampIps[nextId] = append(ampIps[nextId], output)
			if x == true {
				break
			}
		}

		// fmt.Println(ampIps)

		if maxOutput < ampIps[0][len(ampIps[0])-1] {
			maxOutput = ampIps[0][len(ampIps[0])-1]
		}
	}
	fmt.Println(maxOutput)
}
