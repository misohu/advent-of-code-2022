package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	fileName      = "input.txt"
	elfs_in_group = 3
)

func getLetterValue(letter int) int {
	if letter > int('a') {
		return letter - int('a') + 1
	} else {
		return letter - int('A') + 27
	}
}

func processFile(fileName string) int {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0

	for {
		elfs := []string{}
		end := false
		for i := 0; i < elfs_in_group; i++ {
			if !scanner.Scan() {
				end = true
				break
			}
			elfs = append(elfs, scanner.Text())
		}
		if end {
			break
		}
		target := processElfs(elfs)
		result += getLetterValue(target)
	}
	return result
}

func processElfs(elfs []string) int {
	occurences1 := map[byte]bool{}
	occurences2 := map[byte]bool{}

	for i, _ := range elfs[0] {
		occurences1[elfs[0][i]] = true
	}
	for i, _ := range elfs[1] {
		occurences2[elfs[1][i]] = true
	}
	for _, l := range elfs[2] {
		if occurences1[byte(l)] && occurences2[byte(l)] {
			return int(l)
		}
	}

	return 0
}

func main() {
	fmt.Println(processFile(fileName))
}
