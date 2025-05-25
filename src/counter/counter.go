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

func main() {
	args := os.Args
	var number_list []int = []int{}
	var start, end int

	if len(args[1:]) < 2 {
		start, err1 := strconv.Atoi(args[1])
		if err1 != nil {
			fmt.Println("Error: Invalid input. Please provide a valid integer.", args[1])
			return
		}
		number_list = getRange(0, start)
	} else if len(args[1:]) == 2 {
		start, err1 := strconv.Atoi(args[1])
		end, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			fmt.Println("Error: Invalid input. Please provide two integers.")
			return
		}
		number_list = getRange(start, end)
	} else {
		fmt.Println("Error: Please provide one or two integers only.")	
		return
	}
	fmt.Println(fmt.Sprintf("Counting from %d to %d...", start, end))
	for _, number := range number_list {
		fmt.Println(number)
	}
}
