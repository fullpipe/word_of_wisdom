package puzzle

import (
	"crypto/rand"

	"github.com/sirupsen/logrus"
)

const PuzzleSize = 32

type Puzzle struct {
	Complexity int
	Question   []byte
}

func NewPuzzle(complexity int) Puzzle {
	puzzle := Puzzle{
		Complexity: complexity,
		Question:   make([]byte, PuzzleSize),
	}

	_, err := rand.Read(puzzle.Question)
	if err != nil {
		logrus.Fatal(err)
	}

	return puzzle
}
