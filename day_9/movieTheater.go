package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "test.txt")
	fmt.Println(path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	var myPoints []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && len(scanner.Text()) != 0 {
		myPoints = append(myPoints, buildPoint(scanner.Text()))
	}

	fmt.Println(myPoints)
	fmt.Println(findLargestRectangle(myPoints))

}

func findLargestRectangle(myPoints []Point) int {
	max := 0
	i := 0
	for i < len(myPoints) {
		j := i + 1
		for j < len(myPoints) {
			if validateConners(myPoints, myPoints[i], myPoints[j]) {
				area := findArea(myPoints[i], myPoints[j])
				fmt.Println("Area for", myPoints[i], myPoints[j], area)
				if max < area {
					max = area
				}
			} else {
				fmt.Println("Invalid Conner", myPoints[i], myPoints[j])
			}
			j++
		}
		i++
	}
	return max
}

func validateConners(myPoints []Point, a, b Point) bool {
	c1 := Point{a.X, b.Y}
	c2 := Point{b.X, a.Y}
	fmt.Println("Validating", c1, c2)
	return validatePoint(myPoints, c1) && validatePoint(myPoints, c2)
}

func validatePoint(myPoints []Point, p Point) bool {
	intersections := 0
	var previousPoint Point
	i := 1
	for i < len(myPoints) {
		pointA := myPoints[i-1]
		pointB := myPoints[i]
		// check point is
		//  1. above min y
		//  2. below max y
		//  3. below the max x (to the left)
		// if this is true, a horizontal line will intersect this edge and we can count it

	}

	// handle first/last point

	return (intersections%2 == 1) // odd number of intersections indicate the point is within the polynomial
}

func findArea(a, b Point) int {
	xDist := math.Abs(float64(a.X - b.X))
	yDist := math.Abs(float64(a.Y - b.Y))
	return int(math.Round((xDist + 1) * (yDist + 1)))
}

func buildPoint(line string) Point {
	parts := strings.Split(line, ",")
	// eating errors
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Point{x, y}
}
