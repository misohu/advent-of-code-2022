package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	fileName = "input.txt"
)

type Point struct {
	X, Y int
}

type Head struct {
	*Point
	History *Point
}

func (p *Point) areNeighbours(p2 *Point) bool {
	dX := math.Abs(float64(p.X - p2.X))
	dY := math.Abs(float64(p.Y - p2.Y))

	if dX > 1 {
		return false
	}
	if dY > 1 {
		return false
	}
	return true
}

func (p *Point) toString() string {
	x := strconv.Itoa(p.X)
	y := strconv.Itoa(p.Y)

	return fmt.Sprintf("[%s,%s]", x, y)
}

func (h *Head) move(dir string) {
	h.History = h.Point
	switch dir {
	case "R":
		h.Point = &Point{h.X + 1, h.Y}
	case "L":
		h.Point = &Point{h.X - 1, h.Y}
	case "U":
		h.Point = &Point{h.X, h.Y + 1}
	case "D":
		h.Point = &Point{h.X, h.Y - 1}
	}
}

func moveThings(h *Head, tail *Point, moves []string) map[string]bool {
	tailVisited := map[string]bool{
		"[0,0]": true,
	}
	for _, move := range moves {
		h.move(move)
		if !h.areNeighbours(tail) {
			tail = h.History
			tailVisited[tail.toString()] = true
		}
		fmt.Printf("%s %s %s %v\n", move, h.Point.toString(), tail.toString(), h.areNeighbours(tail))
	}
	return tailVisited
}

func readFile(fileName string) []string {
	res := []string{}

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		num, _ := strconv.Atoi(parts[1])
		for i := 0; i < num; i++ {
			res = append(res, parts[0])
		}
	}
	return res
}

func main() {
	directions := readFile(fileName)

	fmt.Println(directions)
	res := moveThings(&Head{Point: &Point{0, 0}}, &Point{0, 0}, directions)
	fmt.Println(len(res))
}
