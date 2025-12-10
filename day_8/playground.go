package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var MAX_CONNECTIONS = 1000

type JunctionBox struct {
	X, Y, Z, Id int
}

type Connection struct {
	A, B     JunctionBox
	Distance float64
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "my_input.txt")
	fmt.Println(path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	var myBoxes []JunctionBox
	var myCircuits [][]JunctionBox
	scanner := bufio.NewScanner(file)
	i := 0 // arbiratry id to track junctionBox connections
	for scanner.Scan() && len(scanner.Text()) != 0 {
		myBoxes = append(myBoxes, buildJunctionBox(scanner.Text(), i))
		myCircuits = append(myCircuits, []JunctionBox{myBoxes[i]})
		i++
	}

	possibleConnections := buildPossibleConnections(myBoxes)
	sort.Slice(possibleConnections, func(i, j int) bool {
		return possibleConnections[i].Distance < possibleConnections[j].Distance
	})

	i = 0
	for i < len(possibleConnections) {
		myCircuits = connect(possibleConnections[i].A, possibleConnections[i].B, myCircuits)
		i++
	}

	sort.Slice(myCircuits, func(i, j int) bool {
		return len(myCircuits[i]) > len(myCircuits[j])
	})

	// printMyCircuits(myCircuits)
	// fmt.Println(len(myCircuits[0]) * len(myCircuits[1]) * len(myCircuits[2]))
}

func connect(a JunctionBox, b JunctionBox, myCircuits [][]JunctionBox) [][]JunctionBox {
	circuitAIndex := findCircuitForId(myCircuits, a.Id)
	circuitBIndex := findCircuitForId(myCircuits, b.Id)
	if circuitAIndex != circuitBIndex {
		// append b to a
		myCircuits[circuitAIndex] = append(myCircuits[circuitAIndex], myCircuits[circuitBIndex]...)
		// remove b
		myCircuits = append(myCircuits[:circuitBIndex], myCircuits[circuitBIndex+1:]...)
		fmt.Println("Connecting:", a.X*b.X)
	}
	return myCircuits
}

func findCircuitForId(myCircuits [][]JunctionBox, id int) int {
	for index, circuit := range myCircuits {
		for _, jb := range circuit {
			if jb.Id == id {
				return index
			}
		}
	}
	return -1
}

func buildPossibleConnections(myBoxes []JunctionBox) []Connection {
	var connections []Connection
	i := 0
	for i < len(myBoxes) {
		j := i + 1
		for j < len(myBoxes) {
			distance := findDistance(myBoxes[i], myBoxes[j])
			connections = append(connections, Connection{myBoxes[i], myBoxes[j], distance})
			j++
		}
		i++
	}
	return connections

}

func buildJunctionBox(line string, id int) JunctionBox {
	parts := strings.Split(line, ",")
	// eating errors
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return JunctionBox{x, y, z, id}
}

func findDistance(a JunctionBox, b JunctionBox) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2) + math.Pow(float64(a.Z-b.Z), 2))
}

func printMyCircuits(myCircuits [][]JunctionBox) {
	for index, circuit := range myCircuits {
		//for _, jb := range circuit {
		//}
		fmt.Println(index, circuit)
	}

}
