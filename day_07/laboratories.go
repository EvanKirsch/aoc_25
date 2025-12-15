package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var TACYON_START = []byte("S")[0]
var EMPTY_SPACE = []byte(".")[0]
var TACYON_BEAM = []byte("|")[0]
var TACYON_SPLITER = []byte("^")[0]
var STR_TACYON_BEAM = "|"
var SPLITTER_TIMELINES map[Location]int

type Location struct {
	Line, StrIndex int
}

func main() {
	SPLITTER_TIMELINES = make(map[Location]int)
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
	for scanner.Scan() && len(scanner.Text()) != 0 {
		myLines = append(myLines, scanner.Text())
	}

	var startIndex = strings.IndexByte(myLines[0], TACYON_START)
	fireLocation := Location{1, startIndex}
	splitCount, _ := fireTachyonBeam(myLines, fireLocation, 0, 0)
	timelines := _fireTachyonBeam(myLines, fireLocation, 0)
	printLines(myLines)
	fmt.Println("Split Count: ", splitCount)
	fmt.Println("Timelines: ", timelines)

}

func fireTachyonBeam(myLines []string, fireLocation Location, splitCount int, timelines int) (int, int) {
	if fireLocation.Line > (len(myLines)-1) || fireLocation.StrIndex > len(myLines[0]) || fireLocation.StrIndex < 0 {
		return splitCount, (timelines + 1)
	}
	curChar := myLines[fireLocation.Line][fireLocation.StrIndex]
	switch curChar {
	case EMPTY_SPACE:
		myLines[fireLocation.Line] = myLines[fireLocation.Line][:fireLocation.StrIndex] + STR_TACYON_BEAM + myLines[fireLocation.Line][fireLocation.StrIndex+1:]
		splitCount, timelines = fireTachyonBeam(myLines, Location{fireLocation.Line + 1, fireLocation.StrIndex}, splitCount, timelines)
	case TACYON_SPLITER:
		splitCount = splitCount + 1
		splitCount, timelines = fireTachyonBeam(myLines, Location{fireLocation.Line, fireLocation.StrIndex - 1}, splitCount, timelines)
		splitCount, timelines = fireTachyonBeam(myLines, Location{fireLocation.Line, fireLocation.StrIndex + 1}, splitCount, timelines)
	}
	return splitCount, timelines
}

func _fireTachyonBeam(myLines []string, fireLocation Location, timelines int) int {
	if fireLocation.Line > (len(myLines)-1) || fireLocation.StrIndex > len(myLines[0]) || fireLocation.StrIndex < 0 {
		return (timelines + 1)
	}
	curChar := myLines[fireLocation.Line][fireLocation.StrIndex]
	switch curChar {
	case EMPTY_SPACE, TACYON_BEAM:
		timelines = _fireTachyonBeam(myLines, Location{fireLocation.Line + 1, fireLocation.StrIndex}, timelines)
	case TACYON_SPLITER:
		timelines += 1
		if timelineCount, exists := SPLITTER_TIMELINES[fireLocation]; exists {
			return timelineCount
		} else {
			thisLocationTimelines := _fireTachyonBeam(myLines, Location{fireLocation.Line, fireLocation.StrIndex - 1}, 0)
			thisLocationTimelines += _fireTachyonBeam(myLines, Location{fireLocation.Line, fireLocation.StrIndex + 1}, 0)
			SPLITTER_TIMELINES[fireLocation] = thisLocationTimelines
			return thisLocationTimelines
		}
	}
	return timelines
}

func printLines(myLines []string) {
	for _, line := range myLines {
		fmt.Println(line)
	}
}
