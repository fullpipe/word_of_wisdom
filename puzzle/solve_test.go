package puzzle

import (
	"crypto/sha256"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solve(t *testing.T) {
	t.Run("solves simple puzzles", func(t *testing.T) {
		puzzle := NewPuzzle(1)
		solution := Solve(puzzle)

		h := sha256.New()
		h.Write(puzzle.Question)
		h.Write(solution.Answer)

		assert.True(t, strings.HasPrefix(string(h.Sum(nil)), "0"))
		assert.False(t, strings.HasPrefix(string(h.Sum(nil)), "000"))
		assert.GreaterOrEqual(t, solution.Attempts, uint64(1))
	})

	t.Run("solves puzzles", func(t *testing.T) {
		puzzle := NewPuzzle(2)
		solution := Solve(puzzle)

		h := sha256.New()
		h.Write(puzzle.Question)
		h.Write(solution.Answer)

		assert.True(t, strings.HasPrefix(string(h.Sum(nil)), "00"))
		assert.GreaterOrEqual(t, solution.Attempts, uint64(1))
	})

	t.Run("solves random puzzle", func(t *testing.T) {
		puzzle := NewPuzzle(3)
		solution := Solve(puzzle)

		h := sha256.New()
		h.Write(puzzle.Question)
		h.Write(solution.Answer)

		assert.True(t, strings.HasPrefix(string(h.Sum(nil)), "000"))
		assert.GreaterOrEqual(t, solution.Attempts, uint64(1))
	})
}
