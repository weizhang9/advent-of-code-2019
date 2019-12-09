package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/weizhang9/advent-of-code-2019/utils"
)

func main() {
	input := utils.GetFileContent("./input.txt")
	wires := strings.Split(string(input), "\n")

	wireA := strings.Split(wires[0], ",")
	wireB := strings.Split(wires[1], ",")

	a := getFootsteps(wireA)
	b := getFootsteps(wireB)

	crossingsA1, crossingsB1 := getCrossings(a, b)
	crossingsA1 = append(crossingsA1, crossingsB1...)

	crossingsA2, crossingsB2 := getCrossings(a, b)

	fmt.Println("Day 3 Part One - Manhattan Distance is:", getManhattanDistance(crossingsA1))
	fmt.Println("Day 3 Part Two - Lowest combined step is:", getLowestCombinedSteps(crossingsA2, crossingsB2))

}

type axis struct {
	x int
	y int
	steps int
}

func move(coordinates *axis, direction string, steps int) *axis {
	switch direction {
	case "L":
		coordinates.x--
	case "R":
		coordinates.x++
	case "D":
		coordinates.y--
	case "U":
		coordinates.y++
	}

	coordinates.steps++

	return coordinates
}

func getInstruction(instruction string) (string, int) {
	direction := string(instruction[0])
	steps, err := strconv.Atoi(instruction[1:])
	utils.CheckError(err, "ERROR: can't convert steps to int")

	return direction, steps
}

func getFootsteps(wireInstructions []string) []axis {
	coordinates := new(axis)
	wireTrace := make([]axis, 0)
	for _, v := range wireInstructions {
		d, s := getInstruction(v)
		for i := 0; i < s; i++ {
			pos := move(coordinates, d, s)
			wireTrace = append(wireTrace, *pos)
		}
	}

	return wireTrace
}

func contains(comparingSet []axis, comparedItem axis) axis {
	for _, v := range comparingSet {
		if v.x == comparedItem.x && v.y == comparedItem.y {
			return v
		}
	}

	return *new(axis)
}

func getManhattanDistance(crossings []axis) int {
	distances := make([]int, 0)
	for _, v := range crossings {
		distance := int(math.Abs(float64(v.x))) + int(math.Abs(float64(v.y)))
		distances = append(distances, distance)
	}
	sort.Ints(distances)

	return distances[0]
}

func getCrossings(a, b []axis) (crossingsA, crossingsB []axis){
	crossingsA = make([]axis, 0)
	crossingsB = make([]axis, 0)

	for _, v := range a {
		bv := contains(b, v) 
		if bv != (axis{}) {
			crossingsA = append(crossingsA, v)
			crossingsB = append(crossingsB, bv)
		}
	}
	return;
}

func getLowestCombinedSteps(a, b []axis) int {
	steps := make([]int, 0)
	for i := 0; i < len(a); i++ {
		if a[i].x == b[i].x && a[i].y == b[i].y {
			steps = append(steps, a[i].steps + b[i].steps)
			steps = utils.UniqueInts(steps)
		}
		
	}
	sort.Ints(steps)

	return steps[0]
}