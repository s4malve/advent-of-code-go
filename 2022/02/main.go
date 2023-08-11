package main

import (
	"fmt"
	"strings"

	"github.com/s4malve/advent-of-code-go/utils"
)

const (
	year = "2022"
	day  = "02"
)

var (
	OPONENT_LETTERS = map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}
	YOUR_LETTERS = map[string]string{
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}
	ROUND_LETTERS = map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
	}
)

var (
	LOOSER_SHAPES = map[string]string{
		"rock":     "scissors",
		"scissors": "paper",
		"paper":    "rock",
	}
	WINNER_SHAPES = map[string]string{
		"scissors": "rock",
		"paper":    "scissors",
		"rock":     "paper",
	}
)

var (
	SHAPE_SCORE = map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}
	ROUND_SCORE = map[string]int{
		"win":  6,
		"draw": 3,
		"lost": 0,
	}
)

func main() {
	part := utils.ParsePartFlag()

	if part == 1 {
		partOne()
	} else {
		partTwo()
	}

}

func getRoundState(opponentLetter, yourLetter string) (roundState string) {
	opponentShape := OPONENT_LETTERS[opponentLetter]
	yourShape := YOUR_LETTERS[yourLetter]
	winState := (yourShape == "rock" && opponentShape == "scissors") ||
		(yourShape == "scissors" && opponentShape == "paper") ||
		(yourShape == "paper" && opponentShape == "rock")

	if opponentShape == yourShape {
		return "draw"
	} else if winState {
		return "win"
	} else {
		return "lost"
	}
}

func getRoundScore(opponentLetter, yourLetter string) int {
	roundState := getRoundState(opponentLetter, yourLetter)
	yourShape := YOUR_LETTERS[yourLetter]

	shapeScore := SHAPE_SCORE[yourShape]
	roundScore := ROUND_SCORE[roundState]

	return shapeScore + roundScore
}

func partOne() {
	f, fileScanner := utils.GetInputFileScanner(year, day)
	defer f.Close()

	totalScore := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		round := strings.Split(txt, " ")
		oponentLetter := round[0]
		yourLetter := round[1]
		roundScore := getRoundScore(oponentLetter, yourLetter)
		totalScore += roundScore
	}

	utils.PrintAdventResult(utils.AdventResult{
		Year:    year,
		Day:     day,
		Message: fmt.Sprintf("the total score is %d", totalScore),
		Part:    1,
	})
}

func partTwo() {
}
