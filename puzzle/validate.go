package puzzle

import (
	"crypto/rand"
	"log"
	"strings"

	"golang.org/x/crypto/scrypt"
)

type Validator struct {
	salt []byte
}

func NewValidator() *Validator {
	v := &Validator{}

	rand.Read(v.salt)

	return v
}

func (v *Validator) Validate(p Puzzle, s Solution) bool {
	dk, err := scrypt.Key(append(p.Question[:], s.Answer[:]...), v.salt, 4, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}

	return strings.HasPrefix(string(dk), strings.Repeat("0", p.Complexity))
}
