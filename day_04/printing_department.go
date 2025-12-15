package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var ROLL_OF_PAPER = byte(64)
var X_ROLL_OF_PAPER = byte(88)
var lineLen int

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
	scanner := bufio.NewScanner(file)
	var myMapBuilder strings.Builder
	for scanner.Scan() {
		lineLen = len(scanner.Text())
		fmt.Fprintf(&myMapBuilder, scanner.Text())
	}

	myMap := myMapBuilder.String()
	removedCount := 0
	accessableCount := 1
	for accessableCount != 0 {
		accessableCount, myMap = idRemoveableRolls(myMap)
		fmt.Println("")
		fmt.Print("Removing ")
		fmt.Print(accessableCount)
		fmt.Println(" Rolls of Paper")
		fmt.Println("")
		printMap(myMap)
		myMap = removeRolls(myMap)
		removedCount += accessableCount
	}

	fmt.Println("")
	fmt.Println(removedCount)
}

func idRemoveableRolls(myMap string) (int, string) {
	i := 0
	accessableCount := 0
	for i < len(myMap) {
		if myMap[i] == ROLL_OF_PAPER && getAdjacentCount(i, myMap) < 4 {
			accessableCount++
			myMap = myMap[:i] + "X" + myMap[i+1:]
		}
		i++
	}
	return accessableCount, myMap
}

func removeRolls(myMap string) string {
	i := 0
	for i < len(myMap) {
		if myMap[i] == X_ROLL_OF_PAPER {
			myMap = myMap[:i] + "." + myMap[i+1:]
		}
		i++
	}
	return myMap
}

func getAdjacentCount(index int, myMap string) int {
	neighbors := getNeighborIndexes(index)
	count := 0
	i := 0
	for i < len(neighbors) {
		n := neighbors[i]
		if n > 0 && n < len(myMap) {
			myChar := myMap[n]
			if myChar == ROLL_OF_PAPER || myChar == X_ROLL_OF_PAPER {
				count++
			}
		}
		i++
	}
	return count
}

func getNeighborIndexes(index int) [8]int {
	myIndexes := [8]int{
		index - lineLen - 1,
		index - lineLen,
		index - lineLen + 1,
		index - 1,
		index + 1,
		index + lineLen - 1,
		index + lineLen,
		index + lineLen + 1,
	}
	if index%lineLen == 0 {
		myIndexes[0] = -1
		myIndexes[3] = -1
		myIndexes[5] = -1
	}
	if index%lineLen == lineLen-1 {
		myIndexes[2] = -1
		myIndexes[4] = -1
		myIndexes[7] = -1
	}
	return myIndexes
}

func printMap(myMap string) {
	i := 0
	for i < len(myMap) {
		if (i%lineLen) == 0 && i != 0 {
			fmt.Println("")
		}
		fmt.Print(string(myMap[i]))
		i++
	}
	fmt.Println("")
}
