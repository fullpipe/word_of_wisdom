package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solve(t *testing.T) {
	t.Run("solves puzzles", func(t *testing.T) {
		puzzle := NewPuzzle(1)
		solution := NewSolver().Solve(puzzle)

		assert.GreaterOrEqual(t, solution.Attempts, uint64(1))
	})
}
