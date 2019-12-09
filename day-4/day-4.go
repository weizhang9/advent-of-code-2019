package main

import (
	"fmt"
	"strconv"

	"github.com/weizhang9/advent-of-code-2019/utils"
)

func main() {
	c := 0
	for num := 158126; num < 624574; num++ {
		numSlice := numsToIntSlice(num)
		if pass(numSlice) {
			c++
		}
	}

	fmt.Println("Day 4 Part One:", c)
}

func numsToIntSlice(num int) []int {
	numstr := strconv.Itoa(num)
	ints:= make([]int, 0, len(numstr))
	for _, v := range numstr {
		n, err := strconv.Atoi(string(v))
		utils.CheckError(err, "Can't convert string to int")
		ints = append(ints, n)
	}
	return ints
}

func pass(ints []int) bool {
	hasSameDigits := false
	for i := 0; i < len(ints)-1; i++ {
		if ints[i] > ints[i+1] {
			return false
		}

		if ints[i] == ints[i+1] {
			hasSameDigits = true
		}
	}

	return hasSameDigits
}
