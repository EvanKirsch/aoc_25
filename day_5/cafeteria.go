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

type MyRange struct {
	Min, Max int
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
	scanner := bufio.NewScanner(file)
	var myRanges []*MyRange
	for scanner.Scan() && len(scanner.Text()) != 0 {
		myRanges = append(myRanges, getRange(scanner.Text()))
	}

	freshCount := getFreshCount(scanner, myRanges)
	fmt.Println(freshCount)
	fmt.Println("---")
	listSize := getFreshListSize(myRanges)
	fmt.Println(listSize)

}

func getFreshListSize(myRanges []*MyRange) int {
	myRanges = simplifyRanges(myRanges)
	return getTotalInRanges(myRanges)
}

func getTotalInRanges(myRanges []*MyRange) int {
	runningTotal := 0
	i := 0
	for i < len(myRanges) {
		runningTotal += (myRanges[i].Max - myRanges[i].Min) + 1 // adds 1 for inclusivity
		i++
	}
	return runningTotal
}

func simplifyRanges(myRanges []*MyRange) []*MyRange {
	sort.Slice(myRanges, func(i, j int) bool {
		return myRanges[i].Min < myRanges[j].Min
	})
	i := 0
	for i < len(myRanges) {
		if i+1 < len(myRanges) && myRanges[i].Max >= myRanges[i+1].Min {
			minValue := int(math.Min(float64(myRanges[i].Min), float64(myRanges[i+1].Min)))
			maxValue := int(math.Max(float64(myRanges[i].Max), float64(myRanges[i+1].Max)))
			myRanges[i+1] = &MyRange{minValue, maxValue}
			myRanges = append(myRanges[:i], myRanges[i+1:]...) // remove combined item
			i--
		}
		i++
	}

	return myRanges
}

func printMyRanges(myRanges []*MyRange) {
	i := 0
	for i < len(myRanges) {
		fmt.Println(myRanges[i].Min, myRanges[i].Max)
		i++
	}
}

func getFreshCount(scanner *bufio.Scanner, myRanges []*MyRange) int {
	freshCount := 0
	for scanner.Scan() && len(scanner.Text()) != 0 {
		if checkRanges(myRanges, scanner.Text()) {
			freshCount++
		}
	}
	return freshCount
}

func getRange(myString string) *MyRange {
	myStrings := strings.Split(myString, "-")
	i0, err := strconv.Atoi(myStrings[0])
	if err != nil {
		fmt.Println(err)
	}

	i1, err := strconv.Atoi(myStrings[1])
	if err != nil {
		fmt.Println(err)
	}

	myRange := MyRange{i0, i1}
	return &myRange
}

func checkRanges(myRanges []*MyRange, strValue string) bool {
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		fmt.Println(err)
	}

	i := 0
	for i < len(myRanges) {
		if intValue >= myRanges[i].Min && intValue <= myRanges[i].Max {
			return true
		}
		i++
	}
	return false
}
