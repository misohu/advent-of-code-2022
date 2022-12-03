package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	fileName = "input.txt"
)

func processFile(fileName string) int {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		target := processRucksack(rucksack)
		res += getLetterValue(target)
	}

	return res
}

func processRucksack(rucksack string) int {
	occurences := map[byte]bool{}
	b := len(rucksack) / 2

	for i, _ := range rucksack[:b] {
		occurences[rucksack[i]] = true
	}

	for _, l := range rucksack[b:] {
		if occurences[byte(l)] {
			return int(l)
		}
	}

	return ' '
}

func getLetterValue(letter int) int {
	if letter > int('a') {
		return letter - int('a') + 1
	} else {
		return letter - int('A') + 27
	}
}

func main() {
	fmt.Println(processFile(fileName))
}
