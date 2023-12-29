package puzzle

import (
	"crypto/rand"

	"github.com/sirupsen/logrus"
)

func Solve(p Puzzle) Solution {
	solution := Solution{}

	for {
		solution.Attempts++
		solution.Answer = make([]byte, AnswerSize)
		_, err := rand.Read(solution.Answer)
		if err != nil {
			logrus.Fatal(err)
		}

		if Validate(p, solution) {
			return solution
		}
	}
}
