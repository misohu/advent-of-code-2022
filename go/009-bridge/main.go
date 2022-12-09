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
	X, Y    int
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

func (p *Point) move(dir string) *Point {

	switch dir {
	case "R":
		p = &Point{p.X + 1, p.Y, p}
	case "L":
		p = &Point{p.X - 1, p.Y, p}
	case "U":
		p = &Point{p.X, p.Y + 1, p}
	case "D":
		p = &Point{p.X, p.Y - 1, p}
	}
	return p
}

func moveThings(h *Point, tails []*Point, moves []string) map[string]bool {
	tailVisited := map[string]bool{
		"[0,0]": true,
	}
	for _, move := range moves {
		h = h.move(move)
		if !h.areNeighbours(tails[0]) {
			tails[0] = h.History
			// tails[0] = h.History
		}

		for i := 1; i < len(tails); i++ {
			if !tails[i].areNeighbours(tails[i-1]) {
				// tails[i].History = tails[i]
				tails[i] = tails[i-1].History
				// tails[i] = &Point{tails[i-1].History.X, tails[i-1].History.Y, tails[i]}
				// if i == 8 {
				// 	fmt.Println(tails[i].toString())
				// }
			}
		}
		tailVisited[tails[8].toString()] = true
		// fmt.Printf("%s %s %s %v\n", move, h.toString(), tails[0].toString(), tailVisited)
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
	tails := []*Point{}
	start := &Point{0, 0, nil}
	for i := 0; i < 10; i++ {
		tails = append(tails, start)
	}
	res := moveThings(tails[0], tails[1:], directions)
	fmt.Println(res)
	fmt.Println(len(res))
}
