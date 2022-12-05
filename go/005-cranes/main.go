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
	numRows  = 9
)

type Move struct {
	num  int
	from string
	to   string
}

func processFile(fileName string) (map[string][]string, []Move) {
	file, _ := os.Open(fileName)
	defer file.Close()

	rows := map[string][]string{}
	moves := make([]Move, 0)
	scanner := bufio.NewScanner(file)
	for i := 0; i < numRows; i++ {
		rows[strconv.Itoa(i)] = make([]string, 0)
	}

	for scanner.Scan() {
		row := scanner.Text()

		if string(row[1]) == "1" {
			break
		}

		craneI := 0
		for i := 1; i < len(row); i += 4 {
			c := string(row[i])
			if c != " " {
				rows[strconv.Itoa(craneI)] = append(rows[strconv.Itoa(craneI)], c)
			}
			craneI++
		}
	}

	scanner.Scan()
	for scanner.Scan() {
		row := scanner.Text()

		elements := strings.Split(row, " ")
		num, _ := strconv.Atoi(elements[1])
		from, _ := strconv.Atoi(elements[3])
		to, _ := strconv.Atoi(elements[5])
		moves = append(moves, Move{
			num,
			strconv.Itoa(from - 1),
			strconv.Itoa(to - 1),
		})
	}

	return rows, moves
}

func processData(start map[string][]string, moves []Move) map[string][]string {
	for _, move := range moves {
		fmt.Printf("%v %v -> ", start, move)
		newFrom := make([]string, len(start[move.from]))
		// newTo := make([]string, len(start[move.to]))
		copy(newFrom, start[move.from])
		// copy(newTo, start[move.to])
		// for i := 0; i < move.num; i++ {
		// 	start[move.to] = append([]string{newFrom[i]}, start[move.to]...)
		// }
		start[move.to] = append(newFrom[:move.num], start[move.to]...)
		start[move.from] = start[move.from][move.num:]
		fmt.Println(start)
	}

	return start
}

func main() {
	start, moves := processFile(fileName)
	fmt.Println(start)
	start = processData(start, moves)
	fmt.Println(start)
	res := []string{}
	for i := 0; i < numRows; i++ {
		index := strconv.Itoa(i)
		res = append(res, start[index][0])
	}
	fmt.Println(res)
}
