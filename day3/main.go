package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func readfile(filename string) [][]byte {
	grid := make([][]byte, 0)
	file, _ := os.ReadFile(filename)
	data := string(file)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		grid_line := make([]byte, 0)
		for _, char := range line {
			grid_line = append(grid_line, byte(char))
		}
		grid = append(grid, grid_line)
	}
	return grid

}
func search_symbol_coord(grid [][]byte) []coord {
	symbol_coord_list := make([]coord, 0)
	for i, line := range grid {
		for j, char := range line {
			if (char < '0' || char > '9') && char != '.' {
				symbol_coord_list = append(symbol_coord_list, coord{i, j})
			}
		}
	}
	return symbol_coord_list
}
func search_number_near(symbols []coord, grid [][]byte) []coord {
	number_coord_list := make([]coord, 0)
	seen := map[coord]bool{}
	directions := [][2]int{
		{-1, 0},  // Haut
		{1, 0},   // Bas
		{0, -1},  // Gauche
		{0, 1},   // Droite
		{-1, -1}, // Diagonale haut-gauche
		{-1, 1},  // Diagonale haut-droite
		{1, -1},  // Diagonale bas-gauche
		{1, 1},   // Diagonale bas-droite
	}
	for _, symbol := range symbols {
		for _, dir := range directions {
			x := symbol.x + dir[0]
			y := symbol.y + dir[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid) {
				if grid[x][y] > '0' && grid[x][y] < '9' {
					if _, exists := seen[coord{x, y}]; !exists {
						seen[coord{x, y}] = true
						seen[coord{x, y + 1}] = true
						seen[coord{x, y - 1}] = true
						number_coord_list = append(number_coord_list, coord{x, y})
					}

				}
			}
		}
	}
	return number_coord_list
}

func find_full_numbers(part_of_numbers []coord, grid [][]byte) []int {
	numbers := make([]int, 0)
	for _, part_number := range part_of_numbers {
		condition := true
		after_condition := true
		before_condition := true
		i := 1
		string_number := string(grid[part_number.x][part_number.y])
		x := part_number.x
		y := part_number.y
		fmt.Println("number : ", part_number)
		for condition {
			y_before := y - i
			y_after := y + i
			fmt.Println(" y before", y_before, "y_after", y_after)

			if before_condition && y_before >= 0 && y_before < len(grid) {
				fmt.Println(" y before apres if", y_before)
				if grid[x][y_before] <= '9' && grid[x][y_before] >= '0' {
					fmt.Println(" y before avant ajout", y_before)
					string_number = string(grid[x][y_before]) + string_number
				} else {
					before_condition = false
				}
			} else {
				before_condition = false
			}

			if after_condition && y_after >= 0 && y_after < len(grid) {
				fmt.Println(" y after apres if", y_after)
				if grid[x][y_after] <= '9' && grid[x][y_after] >= '0' {
					fmt.Println(" y after avant ajout", y_after)
					string_number = string_number + string(grid[x][y_after])
				} else {
					after_condition = false
				}
			} else {
				after_condition = false
			}

			i++
			if !after_condition && !before_condition {
				condition = false
			}
		}
		int_value, _ := strconv.Atoi(string_number)
		numbers = append(numbers, int_value)
	}
	return numbers
}

func sum_tab(tab []int) int {
	sum := 0
	for _, element := range tab {
		sum += element
	}
	return sum
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : go run main.go <filename>")
		os.Exit(1)
	}
	grid := readfile(os.Args[1])
	for _, line := range grid {
		fmt.Printf("%c \n", line)
	}
	symbol := search_symbol_coord(grid)
	numbers := search_number_near(symbol, grid)
	fmt.Println(numbers)
	full_numbers := find_full_numbers(numbers, grid)
	fmt.Println("full numbers :", full_numbers)
	sum := sum_tab(full_numbers)
	fmt.Println("Sum :", sum)

}
