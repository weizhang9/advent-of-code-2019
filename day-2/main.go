package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fname := "./input.txt"
	b, err := getFileContent(fname)
	checkErr(err, "Reading file error")

	input := strings.Split(string(b), ",")
	nums1 := stringToIntSlice(input)
	nums2 := stringToIntSlice(input)

	getResultFromParams(nums1, 12, 2)

	getParamsFromResult(nums2, 19690720)

}

func checkErr(e error, info string) {
	if e != nil {
		log.Fatalln(info, e)
	}
}

func getFileContent(fname string) ([]byte, error) {
	return ioutil.ReadFile(fname)
}

func stringToIntSlice(s []string) []int {
	ints := make([]int, 0, len(s))

	for _, v := range s {
		if len(v) == 0 {
			continue
		}

		i, err := strconv.Atoi(v)
		checkErr(err, "Convert string slice to int slice error")
		ints = append(ints, i)
	}

	return ints
}

func getResultFromParams(nums []int, param1 int, param2 int) {
	nums[1], nums[2] = param1, param2

	for p := range nums {
		rewriteFromFormula(nums, p)
	}
	fmt.Println("Day 2 Part One result", nums[0])
}

func getParamsFromResult(slice []int, result int) {
	nums := make([]int, len(slice))

	for param1 := 0; param1 < 100; param1++ {
		for param2 := 0; param2 < 100; param2++ {
			copy(nums, slice)
			nums[1], nums[2] = param1, param2

			for p := range nums {
				rewriteFromFormula(nums, p)
			}

			if nums[0] == result {
				fmt.Printf("Param1: %d, Param2: %d\n", param1, param2)
				fmt.Println("Day 2 Part Two result", 100*param1+param2)
				return
			}
		}

	}

}

func rewriteFromFormula(nums []int, pos int) {

	if pos%4 == 0 && pos <= len(nums)-4 {
		switch nums[pos] {
		case 1:
			nums[nums[pos+3]] = nums[nums[pos+1]] + nums[nums[pos+2]]
		case 2:
			nums[nums[pos+3]] = nums[nums[pos+1]] * nums[nums[pos+2]]
		case 99:
			break
		}
	}

}
