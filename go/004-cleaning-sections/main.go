package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	fileName = "input.txt"
)

type Section struct {
	Start int
	Stop  int
}

type Assignment struct {
	section1 Section
	section2 Section
}

func (a *Assignment) isSubset() bool {
	if a.section1.Start >= a.section2.Start && a.section1.Stop <= a.section2.Stop {
		return true
	}

	if a.section2.Start >= a.section1.Start && a.section2.Stop <= a.section1.Stop {
		return true
	}

	return false
}

func isOverlap(s1, s2 Section) bool {
	if s1.Start <= s2.Stop && s1.Stop >= s2.Stop {
		return true
	}

	if s1.Start <= s2.Start && s1.Stop >= s2.Start {
		return true
	}

	return false
}

func NewSection(in string) Section {
	// fmt.Println(in)
	borders := strings.Split(in, "-")
	start, _ := strconv.Atoi(borders[0])
	stop, _ := strconv.Atoi(borders[1])
	return Section{start, stop}
}

func readFile(fileName string) []Assignment {
	file, _ := os.Open(fileName)
	defer file.Close()

	result := []Assignment{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		result = append(result, Assignment{
			NewSection(parts[0]),
			NewSection(parts[1]),
		})
	}

	return result
}

func findSubsets(in []Assignment) int {
	result := 0
	for _, a := range in {
		if a.isSubset() {
			fmt.Println(a)
			// result++
		}
	}
	return result
}

func findOverlaps(in []Assignment) int {
	result := 0
	for _, a := range in {
		if isOverlap(a.section1, a.section2) || isOverlap(a.section2, a.section1) {
			// fmt.Println(a)
			result++
		}
	}
	return result
}

func main() {
	ass := readFile(fileName)
	// fmt.Println(ass)
	findSubsets(ass)
	// fmt.Println(findSubsets(ass))
	// fmt.Println(findOverlaps(ass))
}
