package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readfile(filename string) [][]int {
	file, _ := os.ReadFile(filename)
	content := string(file)
	lines := strings.Split(content, "\n")
	grid := make([][]int, 0) // Les tableaux vont 2 à 2 : indice i = played indice i+1 : winning
	for _, line := range lines {
		intermediate_tab := make([][]int, 0)
		played_tab := make([]int, 0)
		winning_tab := make([]int, 0)
		numbers := strings.Split(line, ":")[1]
		splitted_numbers := strings.Split(numbers, "|")
		played_numbers := strings.Fields(splitted_numbers[0])
		winning_numbers := strings.Fields(splitted_numbers[1])
		fmt.Println("Played numbers : ", played_numbers)
		fmt.Println("winning numbers : ", winning_numbers)
		for _, played_number := range played_numbers {
			if played_number == " " {
				fmt.Println(" played vide")
			} else {
				int_value, _ := strconv.Atoi(played_number)
				played_tab = append(played_tab, int_value)
			}
		}
		for _, winning_number := range winning_numbers {
			if winning_number == " " {
				fmt.Println(" played vide")
			} else {
				int_value, _ := strconv.Atoi(winning_number)
				winning_tab = append(winning_tab, int_value)
			}

		}
		intermediate_tab = append(intermediate_tab, played_tab)
		intermediate_tab = append(intermediate_tab, winning_tab)
		grid = append(grid, intermediate_tab...)

	}
	return grid
}
func seach_winning_numbers(grid [][]int) []int {
	score_tab := make([]int, 0)
	for i := 0; i+1 < len(grid); i = i + 2 {
		played_numbers := grid[i]
		winning_numbers := grid[i+1]
		score := 0
		for _, played_number := range played_numbers {
			for _, winning_number := range winning_numbers {
				if played_number == winning_number {
					if score == 0 {
						score += 1
					} else {
						score = score * 2
					}
				}
			}
		}
		score_tab = append(score_tab, score)
	}
	return score_tab
}

func sum_tab(numbers []int) int {
	score := 0
	for _, element := range numbers {
		score += element
	}
	return score
}

/* -------------------------------------------------PART 2 ---------------------------------------------*/
/*
func seach_winning_numbers_part2(grid [][]int) []int {
	score_tab := make([]int, len(grid)/2)
	score_tab = initalise_to_1(score_tab)
	fmt.Println("score tab initialisation :", score_tab)
	for i := 0; i+1 < len(grid); i = i + 2 {
		played_numbers := grid[i]
		winning_numbers := grid[i+1]
		score := 0
		fmt.Println("i: ", (i/2)+1)
		k := 0
		for k < score_tab[(i/2)] {
			fmt.Println("Nombre de boucle : ", k, "score tab actuel :", score_tab)
			for _, played_number := range played_numbers {
				for _, winning_number := range winning_numbers {
					if played_number == winning_number {
						score += 1
					}
				}
			}
			j := 1
			for j < score {
				if ((i / 2) + j) < len(grid)/2 {
					score_tab[(i/2)+j] += 1
				} else {
					break
				}
				j++
			}
			k++
		}

	}
	return score_tab
}
*/

func seach_winning_numbers_part2(grid [][]int) []int {
	// Each card has at least one instance initially
	total_cards := len(grid) / 2
	score_tab := make([]int, total_cards)
	for i := range score_tab {
		score_tab[i] = 1
	}

	// Process each scratchcard
	for i := 0; i < total_cards; i++ {
		played_numbers := grid[i*2]
		winning_numbers := grid[i*2+1]
		match_count := 0

		// Count how many matching numbers exist
		for _, played_number := range played_numbers {
			for _, winning_number := range winning_numbers {
				if played_number == winning_number {
					match_count++
				}
			}
		}

		// Each instance of card i wins `match_count` copies of the next `match_count` cards
		for j := 1; j <= match_count; j++ {
			if i+j < total_cards { // Ensure we don't go out of bounds
				score_tab[i+j] += score_tab[i] // Multiply by existing instances of this card
			}
		}
	}

	return score_tab
}

func initalise_to_1(tab []int) []int {
	for i := 0; i < len(tab); i++ {
		tab[i] = 1
	}
	return tab
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(" Usage :  go run main.go <filename>")
		os.Exit(1)
	}
	grid := readfile(os.Args[1])
	for _, line := range grid {
		fmt.Println(line)
	}
	numbers := seach_winning_numbers(grid)
	fmt.Println("Numbers : ", numbers)
	score := sum_tab(numbers)
	fmt.Println("Score : ", score)
	fmt.Println()
	fmt.Println("-------------------PART 2 --------------------")
	fmt.Println()
	numbers_part2 := seach_winning_numbers_part2(grid)
	fmt.Println("Numbers : ", numbers_part2)
	score_total := sum_tab(numbers_part2)
	fmt.Println("Score : ", score_total)

}
