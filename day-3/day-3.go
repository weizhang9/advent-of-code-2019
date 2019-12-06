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

	crossings := make([]axis, 0)

	for _, v := range a {
		if contains(b, v) {
			crossings = append(crossings, v)
		}
	}

	fmt.Println("Day 3 Part One - Manhattan Distance is:", getManhattanDistance(crossings))

}

type axis struct {
	x int
	y int
}

func move(coordinates *axis, direction string, steps int) *axis {
	switch direction {
	case "L":
		coordinates.x --
	case "R":
		coordinates.x ++
	case "D":
		coordinates.y --
	case "U":
		coordinates.y ++
	}

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

func contains(comparingSet []axis, comparedItem axis) bool {
	for _, v := range comparingSet {
		if v == comparedItem {
			return true
		}
	}

	return false
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
