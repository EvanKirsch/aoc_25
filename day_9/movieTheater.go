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

type Point struct {
	X, Y int
}

type Rectangle struct {
	A, B     Point
	Area, Id int
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

	var myPoints []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && len(scanner.Text()) != 0 {
		myPoints = append(myPoints, buildPoint(scanner.Text()))
	}

	fmt.Println(myPoints)
	rectangles := findRectangles(myPoints)
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].Area < rectangles[j].Area
	})
	rectangle := findLargestValid(myPoints, rectangles)
	fmt.Println(rectangle)

}

func findRectangles(myPoints []Point) []Rectangle {
	i := 0
	id := 0
	var rectangles []Rectangle
	for i < len(myPoints) {
		j := i + 1
		for j < len(myPoints) {
			area := findArea(myPoints[i], myPoints[j])
			rectangle := Rectangle{myPoints[i], myPoints[j], area, id}
			rectangles = append(rectangles, rectangle)
			j++
			id++
		}
		i++
	}
	return rectangles
}

func findLargestValid(myPoints []Point, rectangles []Rectangle) Rectangle {
	var invalidRectangles []Rectangle
	var currentLargest Rectangle
	i := 0
	for _, rectangle := range rectangles {
		fmt.Println("Validating:", rectangle, i)
		if (!hasInvalidSubRectangle(invalidRectangles, rectangle)) && validateSquare(myPoints, rectangle.A, rectangle.B) {
			currentLargest = rectangle
		} else {
			invalidRectangles = append(invalidRectangles, rectangle)
		}
		i++
	}
	return currentLargest
}

func hasInvalidSubRectangle(invalidRectangles []Rectangle, rectangle Rectangle) bool {
	for _, badRectangle := range invalidRectangles {
		if isWithin(badRectangle, rectangle) {
			return true
		}
	}
	return false
}

func isWithin(a, b Rectangle) bool {
	return iMin(b.A.X, b.B.X) <= iMin(a.A.X, a.B.X) && iMax(a.A.X, a.B.X) <= iMax(b.A.X, b.B.X) &&
		iMin(b.A.Y, b.B.Y) <= iMin(a.A.Y, a.B.Y) && iMax(a.A.Y, a.B.Y) <= iMax(b.A.Y, b.B.Y)

}

func validateSquare(myPoints []Point, a, b Point) bool {

	min, max := iMin(a.X, b.X), iMax(a.X, b.X)
	for x := max; x >= min; x-- {
		c1 := Point{x, a.Y}
		c2 := Point{x, b.Y}
		if !(validatePoint(myPoints, c1) && validatePoint(myPoints, c2)) {
			return false
		}
	}

	min, max = iMin(a.Y, b.Y), iMax(a.Y, b.Y)
	for y := max; y >= min; y-- {
		c1 := Point{a.X, y}
		c2 := Point{b.X, y}
		if !(validatePoint(myPoints, c1) && validatePoint(myPoints, c2)) {
			return false
		}
	}

	return true
}

func validatePoint(myPoints []Point, p Point) bool {
	intersections := 0
	i := 1

	for i < len(myPoints) {
		pointA := myPoints[i-1]
		pointB := myPoints[i]
		if isOnEdge(pointA, pointB, p) {
			return true
		} else if crossesEdge(pointA, pointB, p) {
			intersections++
		}
		i++
	}

	// handle first/lasts point
	pointA := myPoints[0]
	pointB := myPoints[len(myPoints)-1]
	if isOnEdge(pointA, pointB, p) {
		return true
	} else if crossesEdge(pointA, pointB, p) {
		intersections++
	}

	return (intersections%2 == 1) // odd number of intersections indicate the point is within the polynomial
}

func crossesEdge(pointA, pointB, p Point) bool {
	return (p.Y > iMin(pointA.Y, pointB.Y) && // greater than lowest y
		p.Y <= iMax(pointA.Y, pointB.Y) && // less than max y
		p.X < iMax(pointA.X, pointB.X)) // and to the left of X
}

func isOnEdge(pointA, pointB, p Point) bool {
	minX := iMin(pointA.X, pointB.X)
	maxX := iMax(pointA.X, pointB.X)
	minY := iMin(pointA.Y, pointB.Y)
	maxY := iMax(pointA.Y, pointB.Y)
	if pointA.X == pointB.X {
		return p.X == pointA.X && minY <= p.Y && p.Y <= maxY
	} else {
		return p.Y == pointA.Y && minX <= p.X && p.X <= maxX
	}

}

func iMax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func iMin(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
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

func drawMap(myPoints []Point) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 14; x++ {
			if pointExists(myPoints, y, x) {
				fmt.Print("#")
			} else if (validatePoint(myPoints, Point{x, y})) {
				fmt.Print("V")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func pointExists(myPoints []Point, y int, x int) bool {
	for _, point := range myPoints {
		if point.X == x && point.Y == y {
			return true
		}
	}
	return false
}
