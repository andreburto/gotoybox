package main

import (
	"fmt"
	"os"
	"strconv"
)

func getRange(s int, e int) ([]int) {
	var rangeList []int
	for i := s; i <= e; i++ {
		rangeList = append(rangeList, i)
	}
	return rangeList
}

func getStartEnd(args []string) (int, int) {
	var start int = 0
	var end int = 0
	var err1, err2 error
	if len(args) == 1 {
		end, err1 = strconv.Atoi(args[0])
	} else if len(args) > 1 {
		start, err1 = strconv.Atoi(args[0])
		end, err2 = strconv.Atoi(args[1])
	} else {
		fmt.Println("Error: Please provide only one or two integer arguments.")
		return -1, -1
	}
	if err1 != nil {
		fmt.Println("Error: Invalid input. Please provide a valid integer.", args)
		return -1, -1
	}
	if err2 != nil {
		fmt.Println("Error: Invalid input. Please provide two valid integers.", args)
		return -1, -1
	}
	return start, end
}

func main() {
	args := os.Args

	var start, end int = getStartEnd(args[1:])

	if start == -1 && end == -1 {
		fmt.Println("Error: Unable to parse start and end values from input.")
		return
	}

	if start > end {
		fmt.Println("Error: Start value must be less than or equal to end value.")
		return
	}

	fmt.Printf("Counting from %d to %d...\n", start, end)

	var number_list []int = getRange(start, end)
	for _, number := range number_list {
		fmt.Println(number)
	}
}
