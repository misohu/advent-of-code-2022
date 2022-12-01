package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Calory struct {
	input string
}

type Elf struct {
	backpack []Calory
}

func (c *Calory) value() (int, error) {
	value, err := strconv.Atoi(c.input)
	if err != nil {
		return -1, fmt.Errorf("Calory.value: cannot convert to int %v", err)
	}
	return value, nil
}

func (e *Elf) sumCalories() (int, error) {
	total := 0
	for _, c := range e.backpack {
		cVal, err := c.value()
		if err != nil {
			return 0, fmt.Errorf("Elf.sumCalories: %v", err)
		}
		total += cVal
	}
	return total, nil
}

func processElfs(fileName string) ([]int, error) {
	elf := Elf{make([]Calory, 0)}
	result := make([]int, 0)

	readFile, err := os.Open(fileName)
	if err != nil {
		return result, fmt.Errorf("processElfs: error while reading file %v", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		input := fileScanner.Text()
		if input == "" {
			elfCalories, err := elf.sumCalories()
			if err != nil {
				return nil, fmt.Errorf("processElfs: %v", err)
			}
			result = append(result, elfCalories)
			elf = Elf{make([]Calory, 0)}
			continue
		}
		elf.backpack = append(elf.backpack, Calory{input})
	}

	elfCalories, err := elf.sumCalories()
	if err != nil {
		return nil, fmt.Errorf("processElfs: %v", err)
	}
	result = append(result, elfCalories)

	return result, nil
}

const (
	fileName = "input.txt"
)

func sumTopN(s []int, n int) int {
	result := 0
	for _, e := range s[:n] {
		result += e
	}
	return result
}

func main() {
	result, err := processElfs(fileName)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})
	fmt.Println(result)

	fmt.Println(sumTopN(result, 3))
}
