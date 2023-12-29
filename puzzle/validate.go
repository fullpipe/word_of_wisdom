package puzzle

import (
	"crypto/sha256"
	"strings"
)

func Validate(p Puzzle, s Solution) bool {
	hasher := sha256.New()

	hasher.Write(p.Question)
	hasher.Write(s.Answer)

	return strings.HasPrefix(string(hasher.Sum(nil)), strings.Repeat("0", p.Complexity))
}
