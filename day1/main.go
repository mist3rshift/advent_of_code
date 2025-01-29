package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readfile(filename string) [][]byte {
	file, _ := os.ReadFile(filename)
	data := string(file)
	lines := strings.Split(data, "\n")
	grid := make([][]byte, 0)
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}
	return grid
}
func extract_numbers(grid [][]byte) [][]byte {
	result := make([][]byte, 0)
	for _, line := range grid {
		numbers_in_line := make([]byte, 0)
		for _, char := range line {
			if char >= '0' && char <= '9' {
				numbers_in_line = append(numbers_in_line, char)
			}
		}
		result = append(result, numbers_in_line)
	}
	return result
}

func sum_line_numbers(numbers [][]byte) []int {
	result := make([]int, 0)
	for _, line := range numbers {
		concat := string(line[0]) + string(line[len(line)-1])
		fmt.Println("concat", concat)
		int_value, _ := strconv.Atoi(concat)
		result = append(result, int_value)
	}
	return result
}

func sum_list(list []int) int {
	var total_sum int
	for _, element := range list {
		total_sum += element
	}
	return total_sum
}

/* -------------------------------------PART 2 ---------------------------------------*/
func extract_numbers_and_written_numbers(grid [][]byte) [][]byte {
	result := make([][]byte, 0)
	for _, line := range grid {
		numbers_in_line := make([]byte, 0)
		letter_used := map[int]bool{}
		for i, char := range line {
			if char >= '0' && char <= '9' {
				letter_used[i] = true
				numbers_in_line = append(numbers_in_line, char)
			}
			if i >= 2 && !letter_used[i] && !letter_used[i-1] && !letter_used[i-2] {
				last_three_letters := string(line[i-2]) + string(line[i-1]) + string(line[i])
				switch last_three_letters {
				case "one":
					numbers_in_line = append(numbers_in_line, '1')
					letter_used[i] = true
					letter_used[i-1] = true
					letter_used[i-2] = true
				case "two":
					numbers_in_line = append(numbers_in_line, '2')
					letter_used[i] = true
					letter_used[i-1] = true
					letter_used[i-2] = true
				case "six":
					numbers_in_line = append(numbers_in_line, '6')
					letter_used[i] = true
					letter_used[i-1] = true
					letter_used[i-2] = true
				}
			}
			if i >= 3 && !letter_used[i] && !letter_used[i-1] && !letter_used[i-2] && !letter_used[i-3] {
				last_four_letters := string(line[i-3]) + string(line[i-2]) + string(line[i-1]) + string(line[i])
				switch last_four_letters {
				case "four":
					numbers_in_line = append(numbers_in_line, '4')
					letter_used[i] = true
					letter_used[i-1] = true
					letter_used[i-2] = true
					letter_used[i-3] = true

				case "five":
					numbers_in_line = append(numbers_in_line, '5')
					letter_used[i] = true
					letter_used[i-1] = true
					letter_used[i-2] = true
					letter_used[i-3] = true
				case "nine":
					numbers_in_line = append(numbers_in_line, '9')
					letter_used[i] = true
					letter_used[i-1] = true
					letter_used[i-2] = true
					letter_used[i-3] = true
				}
			}
			if i >= 4 && !letter_used[i] && !letter_used[i-1] && !letter_used[i-2] && !letter_used[i-3] && !letter_used[i-4] {
				last_four_letters := string(line[i-4]) + string(line[i-3]) + string(line[i-2]) + string(line[i-1]) + string(line[i])
				switch last_four_letters {
				case "three":
					numbers_in_line = append(numbers_in_line, '3')
					letter_used[i] = true
					letter_used[i-1] = true
					letter_used[i-2] = true
					letter_used[i-3] = true
					letter_used[i-4] = true

				case "seven":
					numbers_in_line = append(numbers_in_line, '7')
					letter_used[i] = true
					letter_used[i-1] = true
					letter_used[i-2] = true
					letter_used[i-3] = true
					letter_used[i-4] = true
				case "eight":
					numbers_in_line = append(numbers_in_line, '8')
					letter_used[i] = true
					letter_used[i-1] = true
					letter_used[i-2] = true
					letter_used[i-3] = true
					letter_used[i-4] = true
				}
			}
		}
		result = append(result, numbers_in_line)
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : ./main param")
		os.Exit(1)
	}
	grid := readfile(os.Args[1])
	for _, line := range grid {
		fmt.Printf("%c \n", line)
	}
	fmt.Println()
	fmt.Println()
	/*
		numbers := extract_numbers(grid)
		for _, line := range numbers {
			fmt.Printf("%c \n", line)
		}
		sums := sum_line_numbers(numbers)
		fmt.Println(sums)
		fmt.Println("Total sum : ", sum_list(sums))
	*/
	numbers_and_letters := extract_numbers_and_written_numbers(grid)
	for _, line := range numbers_and_letters {
		fmt.Printf("%c \n", line)
	}
	fmt.Println()
	fmt.Println()
	sums := sum_line_numbers(numbers_and_letters)
	fmt.Println(sums)
	fmt.Println("Total sum : ", sum_list(sums))

}
