package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readfile(filename string) []string {
	file, _ := os.ReadFile(filename)
	data := string(file)
	lines := strings.Split(data, "\n")
	result := make([]string, 0)
	for _, line := range lines {
		fmt.Println("Line: ", line)
		result = append(result, line)
	}
	return result
}

func computeStats(grid []string, rules []string) []int {
	result := make([]int, 0)
	for _, line := range grid {
		allValid := true
		line := strings.ReplaceAll(line, ";", ",")
		elements := strings.Split(line, ":")
		combo_number_color := strings.Split(elements[1], ",")
		fmt.Println("Combo _number _color : ", combo_number_color)
		for _, single_combo := range combo_number_color {
			fmt.Println("single_combo", single_combo)
			number_and_color := strings.Split(single_combo, " ")
			number, _ := strconv.Atoi(number_and_color[1])
			color := strings.ReplaceAll(number_and_color[2], " ", "")
			for _, rule := range rules {
				rule_elements := strings.Split(rule, " ")
				rule_number, _ := strconv.Atoi(rule_elements[0])
				fmt.Println("Rule number and color", color, rule_elements[1])
				if rule_elements[1] == color {
					if number > rule_number {
						allValid = false
						break
					}
				}
			}
			// Si une règle est violée, inutile de vérifier le reste
			if !allValid {
				break
			}
		}
		if allValid {
			result = append(result, 1)
		} else {
			result = append(result, 0)
		}

	}
	return result
}
func computing_score(scorelist []int) int {
	score := 0
	for i := 0; i < len(scorelist); i++ {
		if scorelist[i] != 0 {
			score += i + 1
		}
	}
	return score
}

/* ---------------------------------------------------PART 2 ---------------------------------------*/
func computeStatsPart2(grid []string) []int {
	result := make([]int, 0)
	max_color_value := map[string]int{}
	for _, line := range grid {
		max_color_value["red"] = 0
		max_color_value["blue"] = 0
		max_color_value["green"] = 0
		power_cube := 0
		line := strings.ReplaceAll(line, ";", ",")
		elements := strings.Split(line, ":")
		combo_number_color := strings.Split(elements[1], ",")
		fmt.Println("Combo _number _color : ", combo_number_color)
		for _, single_combo := range combo_number_color {
			fmt.Println("single_combo", single_combo)
			number_and_color := strings.Split(single_combo, " ")
			number, _ := strconv.Atoi(number_and_color[1])
			color := strings.ReplaceAll(number_and_color[2], " ", "")
			switch color {
			case "red":
				if number > max_color_value["red"] {
					max_color_value["red"] = number
				}
			case "blue":
				if number > max_color_value["blue"] {
					max_color_value["blue"] = number
				}
			case "green":
				if number > max_color_value["green"] {
					max_color_value["green"] = number
				}
			}
		}
		power_cube = max_color_value["red"] * max_color_value["blue"] * max_color_value["green"]
		result = append(result, power_cube)

	}
	return result
}
func computing_scorePart2(scorelist []int) int {
	score := 0
	for _, element := range scorelist {
		score += element
	}
	return score
}

func main() {
	if len(os.Args) < 2 {
		println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	content := readfile(os.Args[1])
	fmt.Println(content)
	fmt.Println(content[0])
	rules := []string{"12 red", "13 green", "14 blue"}
	score_list := computeStats(content, rules)
	fmt.Println(computing_score(score_list))
	fmt.Println("-----------------------------PART 2 ------------------------")
	fmt.Println()
	score_part2 := computeStatsPart2(content)
	total_part2 := computing_scorePart2(score_part2)
	fmt.Println(total_part2)

}
