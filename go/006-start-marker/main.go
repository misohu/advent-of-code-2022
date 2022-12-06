package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	Buffer []rune
}

func (s *Stack) Add(r rune) {
	for i, t := range s.Buffer {
		if t == r {
			s.Buffer = s.Buffer[i+1:]
		}
	}
	s.Buffer = append(s.Buffer, r)
	if len(s.Buffer) > patternSize {
		s.Buffer = s.Buffer[1:]
	}
	// fmt.Println(string(s.Buffer))
}

func (s *Stack) IsStart() bool {
	return len(s.Buffer) == patternSize
}

func findStart(input string) int {
	stack := Stack{}
	for i, r := range input {
		stack.Add(r)
		if stack.IsStart() {
			return i + 1
		}
	}
	return -1
}

func readFile(fileName string) string {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

const (
	fileName    = "input.txt"
	patternSize = 14
)

func main() {
	input := readFile(fileName)
	fmt.Println(input)
	fmt.Println(findStart(input))
}
