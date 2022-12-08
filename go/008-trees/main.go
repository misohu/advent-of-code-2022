package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func findTrees(input [][]int) int {
// 	res := 0
// 	for r, row := range input {
// 		for c, _ := range row {
// 			if isHighest2(r, c, input, len(input)) {
// 				res++
// 			}
// 		}
// 	}
// 	return res
// }

func findTrees(input [][]int) int {
	res := 0
	for r, row := range input {
		for c, tree := range row {
			score := isHighest2(r, c, input, len(input))
			fmt.Printf("[%d,%d] -> %d = %d\n", r, c, tree, score)
			if score > res {
				res = score
			}
		}
	}
	return res
}

// func isHighest2(i, j int, input [][]int, max int) bool {
// 	directions := map[string]bool{
// 		"r": true,
// 		"l": true,
// 		"t": true,
// 		"d": true,
// 	}
// 	for r := 0; r < max; r++ {
// 		for direction, _ := range directions {
// 			switch direction {
// 			case "r":
// 				if max-r-1 == j {
// 					fmt.Printf("[%d,%d] -> %d\n", i, j, input[i][j])
// 					return true
// 				}
// 				if input[i][max-r-1] >= input[i][j] {
// 					delete(directions, direction)
// 				}
// 			case "l":
// 				if r == j {
// 					fmt.Printf("[%d,%d] -> %d\n", i, j, input[i][j])
// 					return true
// 				}
// 				if input[i][r] >= input[i][j] {
// 					delete(directions, direction)
// 				}
// 			case "d":
// 				if max-r-1 == i {
// 					fmt.Printf("[%d,%d] -> %d\n", i, j, input[i][j])
// 					return true
// 				}
// 				if input[max-r-1][j] >= input[i][j] {
// 					delete(directions, direction)
// 				}
// 			case "t":
// 				if r == i {
// 					fmt.Printf("[%d,%d] -> %d\n", i, j, input[i][j])
// 					return true
// 				}
// 				if input[r][j] >= input[i][j] {
// 					delete(directions, direction)
// 				}
// 			}
// 		}
// 		if len(directions) == 0 {
// 			return false
// 		}
// 	}
// 	return false
// }

func isHighest2(i, j int, input [][]int, max int) int {
	res := 1
	directions := map[string]bool{
		"r": true,
		"l": true,
		"t": true,
		"d": true,
	}
	for r := 1; r < max-1; r++ {
		if len(directions) == 0 {
			return res
		}
		for direction, _ := range directions {
			switch direction {
			case "r":
				if j+r == max {
					res *= (r-1) 
					delete(directions, direction)
					continue
				}
				if input[i][j+r] >= input[i][j] {
					res *= r
					delete(directions, direction)
				}
			case "l":
				if j-r < 0 {
					res *= (r-1)
					delete(directions, direction)
					continue
				}
				if input[i][j-r] >= input[i][j] {
					res *= r
					delete(directions, direction)
				}
			case "d":
				if i+r == max {
					res *= (r-1)
					delete(directions, direction)
					continue
				}
				if input[i+r][j] >= input[i][j] {
					res *= r
					delete(directions, direction)
				}
			case "t":
				if i-r < 0 {
					res *= (r-1)
					delete(directions, direction)
					continue
				}
				if input[i-r][j] >= input[i][j] {
					res *= r
					delete(directions, direction)
				}
			}
		}
	}
	return res
}

func readFile(fileName string) [][]int {
	res := make([][]int, 0)
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")
		tmp := []int{}
		for _, part := range parts {
			value, _ := strconv.Atoi(part)
			tmp = append(tmp, value)
		}
		res = append(res, tmp)
	}

	return res
}

func main() {
	input := readFile("input.txt")
	fmt.Println(findTrees(input))
	isHighest2(3, 2, input, 5)
}
