package puzzle

import (
	"crypto/rand"

	"github.com/sirupsen/logrus"
)

type Solver struct {
	validator *Validator
}

func NewSolver() *Solver {
	return &Solver{
		validator: NewValidator(),
	}
}

func (s *Solver) Solve(p Puzzle) Solution {
	solution := Solution{}

	for {
		solution.Attempts++
		solution.Answer = make([]byte, AnswerSize)
		_, err := rand.Read(solution.Answer)
		if err != nil {
			logrus.Fatal(err)
		}

		if s.validator.Validate(p, solution) {
			return solution
		}
	}
}
