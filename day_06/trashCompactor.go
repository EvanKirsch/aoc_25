package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

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

	var myLines []string
	scanner := bufio.NewScanner(file)
	myStrNumbers := make(map[int][]string)
	for scanner.Scan() && len(scanner.Text()) != 0 {
		myLines = append(myLines, scanner.Text())
	}

	j := 0
	lineLen := getLineLen(myLines[len(myLines)-1])
	for j < len(myLines) {
		i := 0
		for i < len(lineLen) {
			myStrNumbers[i] = append(myStrNumbers[i], myLines[j][lineLen[i][0]:lineLen[i][1]])
			i++
		}
		j++
	}

	printMyMap(myStrNumbers)
	mySubtotals := evaluateMyMap(myStrNumbers)
	fmt.Println(mySubtotals)
	fmt.Println(sum(mySubtotals))

}

func getLineLen(myString string) [][]int {
	pattern := regexp.MustCompile("(\\*|\\+)\\s*")
	matches := pattern.FindAllStringIndex(myString, -1)
	return matches
}

func sum(myInts []int) int {
	myTotal := 0
	i := 0
	for i < len(myInts) {
		myTotal = myTotal + myInts[i]
		i++
	}
	return myTotal
}

func evaluateMyMap(myMap map[int][]string) []int {
	i := 0
	var mySubTotals []int
	for i < len(myMap) {
		myStrFunction := myMap[i][len(myMap[i])-1]
		myIntNumbers := convertStrArrToInt(myMap[i][:len(myMap[i])-1])
		myCepNumbers := cephlipodNumber(myIntNumbers)
		mySubTotals = append(mySubTotals, applyMyFunction(myStrFunction, myCepNumbers))
		i++
	}
	return mySubTotals
}

func cephlipodNumber(myStringArrs [][]string) []int {
	i := 0
	var myStrNumbers []string
	for i < len(myStringArrs) {
		j := 0
		for j < len(myStringArrs[i]) {
			if len(myStrNumbers) <= j {
				myStrNumbers = append(myStrNumbers, "")
			}
			myStrNumbers[j] = myStrNumbers[j] + string(myStringArrs[i][j])
			j++
		}
		i++
	}

	var myInts []int
	i = 0
	for i < len(myStrNumbers) {
		myStr := strings.Trim(myStrNumbers[i], " ")
		if len(myStr) > 0 {
			myInt, err := strconv.Atoi(myStr)
			myInts = append(myInts, myInt)
			if err != nil {
				fmt.Println(err)
			}
		}
		i++
	}

	return myInts
}

func convertStrArrToInt(myStrArr []string) [][]string {
	i := 0
	var myCharArr [][]string
	for i < len(myStrArr) {
		j := 0
		for j < len(myStrArr[i]) {
			myChar := string(myStrArr[i][j])
			if len(myCharArr) <= i {
				myCharArr = append(myCharArr, []string{})
			}
			myCharArr[i] = append(myCharArr[i], myChar)
			j++
		}
		i++
	}
	return myCharArr
}

func applyMyFunction(myStrFunction string, myIntNumbers []int) int {
	i := 0
	myTotal := 0
	for i < len(myIntNumbers) {
		if strings.Fields(myStrFunction)[0] == "*" {
			if myTotal == 0 {
				myTotal = myIntNumbers[i]
			} else if myIntNumbers[i] == 0 {
				myTotal = myTotal
			} else {
				myTotal = myTotal * myIntNumbers[i]
			}

		} else if strings.Fields(myStrFunction)[0] == "+" {
			myTotal = myTotal + myIntNumbers[i]
		}
		i++
	}
	return myTotal

}

func printMyMap(myMap map[int][]string) {
	i := 0
	for i < len(myMap) {
		fmt.Println(myMap[i])
		i++
	}
}
