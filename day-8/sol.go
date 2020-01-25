package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

var inputFile = flag.String("inputFile", "./input", "Input for solver")
var debug = flag.Bool("debug", false, "Set true for debug output")

func main() {
	flag.Parse()
	f, _ := ioutil.ReadFile(*inputFile)
	ip := string(f)

	if *debug {
		fmt.Println("Size of input is ", len(ip))
		fmt.Println("Contents: \n", ip)
	}

	contents := make([]int, 0)

	for i := 0; i < len(ip); i++ {
		val, _ := strconv.Atoi(ip[i : i+1])
		contents = append(contents, val)
	}

	width := 25
	height := 6
	depth := len(contents) / (width * height)
	fmt.Println(contents, depth)

	zeros := make([]int, depth)
	ones := make([]int, depth)
	twos := make([]int, depth)
	var output [6][25]string

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			output[i][j] = "m"
		}
	}

	var minZeroIndex int
	minZeroCount := math.Inf(1)

	for i := 0; i < depth; i++ {
		for j := i * (width * height); j < (i+1)*(width*height); j++ {
			id := j % (width * height)
			if contents[j] == 0 {
				zeros[i]++
				if output[id/width][id%width] == "m" {
					output[id/width][id%width] = "."
				}
			} else if contents[j] == 1 {
				ones[i]++
				if output[id/width][id%width] == "m" {
					output[id/width][id%width] = "X"
				}
			} else if contents[j] == 2 {
				twos[i]++
			}
		}

		if float64(zeros[i]) < minZeroCount {
			minZeroIndex = i
			minZeroCount = float64(zeros[i])
		}
		fmt.Println(i, float64(zeros[i]))
	}

	fmt.Println(ones[minZeroIndex] * twos[minZeroIndex])
	// fmt.Println(output)
	for i := 0; i < 6; i++ {
		fmt.Println(output[i])
	}
}
