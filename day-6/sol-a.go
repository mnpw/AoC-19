package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var inputFile = flag.String("inputFile", "./input", "Input for solver")
var debug = flag.Bool("debug", false, "Set true for debug output")

type parentNode struct {
	parentName   string
	parentOrbits int
}

func main() {
	flag.Parse()
	f, _ := ioutil.ReadFile(*inputFile)
	ip := strings.Split(string(f), "\n")

	if *debug {
		fmt.Println("Size of input is ", len(ip))
		fmt.Println("Contents: \n", ip)
	}

	parent := make(map[string]parentNode)

	graph := make(map[string][]string)

	for i := 0; i < len(ip); i++ {
		//generate orbit graph
		sep := strings.Index(ip[i], ")")
		par := ip[i][:sep]
		chl := ip[i][sep+1:]
		graph[par] = append(graph[par], chl)
		parent[chl] = parentNode{par, 0}
	}

	curr := "COM"
	fmt.Println(graph[curr])

	added := make(map[string]int)
	orbitCount := make(map[string]int)

	var processQueue []string
	processQueue = append(processQueue, curr)
	orbitCount[curr] = 0
	added[curr] = 1
	checksum := 0

	for len(processQueue) != 0 {

		curr := processQueue[0]
		processQueue = processQueue[1:]

		for i := 0; i < len(graph[curr]); i++ {
			if added[graph[curr][i]] != 1 {
				processQueue = append(processQueue, graph[curr][i])
				added[graph[curr][i]] = 1
				orbitCount[graph[curr][i]] = orbitCount[curr] + 1
				checksum += orbitCount[graph[curr][i]]
			}
		}

	}

	fmt.Println(checksum)

}
