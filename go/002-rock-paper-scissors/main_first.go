// package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	input_file = "input.txt"
)

type Game struct {
	rounds   []Round
	encoding Encoding
	scores   Scores
}

type Round struct {
	p1 string
	p2 string
}

type Encoding map[string]string
type Scores map[string]int

var encoding = Encoding{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

var scores = Scores{
	"A": 1,
	"B": 2,
	"C": 3,
}

var wins = []string{"AB", "BC", "CA"}

func (g *Game) PlayGame() int {
	result := 0
	for _, r := range g.rounds {
		// fmt.Println(g.PlayRound(r))
		result += g.PlayRound(r)
	}
	return result
}

func (g *Game) PlayRound(round Round) int {
	elementValue := g.scores[encoding[round.p2]]
	return elementValue + DecideRound(round)
}

func DecideRound(round Round) int {
	realp2 := encoding[round.p2]
	res := 0

	if round.p1 == realp2 {
		res = 3
	}
	tmp := round.p1 + realp2
	for _, variant := range wins {
		if tmp == variant {
			res = 6
		}
	}

	// fmt.Println(round.p1, realp2, res)
	return res
}

func readFile(fileName string) (*Game, error) {
	game := &Game{
		rounds:   []Round{},
		encoding: encoding,
		scores:   scores,
	}

	readFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("processElfs: error while reading file %v", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		input := fileScanner.Text()
		words := strings.Fields(input)
		game.rounds = append(game.rounds, Round{words[0], words[1]})
	}

	return game, nil
}

func main() {
	game, _ := readFile(input_file)
	fmt.Println(game)
	res := game.PlayGame()
	fmt.Println(res)
}
