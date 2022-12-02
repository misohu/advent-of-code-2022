package main

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

var winEncoding = Encoding{
	"A": "B",
	"B": "C",
	"C": "A",
}

var loseEncoding = Encoding {
	"A": "C",
	"B": "A",
	"C": "B",
}

var scores = Scores{
	"A": 1,
	"B": 2,
	"C": 3,
}

func (g *Game) PlayGame() int {
	result := 0
	for _, r := range g.rounds {
		// fmt.Println(g.PlayRound(r))
		result += g.PlayRound(r)
	}
	return result
}

func (g *Game) PlayRound(round Round) int {
	result := round.p2
	played := round.p1
	gameValue := 3
	if result == "X" {
		played = loseEncoding[round.p1]
		gameValue = 0
	}
	if result == "Z" {
		played = winEncoding[round.p1]
		gameValue = 6
	}
	moveValue := g.scores[played]
	return moveValue + gameValue
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
