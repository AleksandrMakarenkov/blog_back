package password

import (
	"crypto/subtle"
	"encoding/base64"
	"golang.org/x/crypto/argon2"
	"strings"
)

type Comparator struct {
}

func NewComparator() *Comparator {
	return &Comparator{}
}

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// Compare compares two strings, where first is string from user input (e.g login form),
// and second is your hash string in special format (see Hash), which is stored somewhere in database.
// It returns true when they are equal, false otherwise.
// Error will be returned ONLY if something went wrong, so if strings are not equal it will return "false" as a result
// and nil as error because everything was fine during comparison process.
func (c *Comparator) Compare(input string, storedHash string) (bool, error) {
	p := &params{
		memory:      32 * 1024,
		iterations:  3,
		parallelism: 4,
		saltLength:  16,
		keyLength:   32,
	}
	// get saved hash and salt
	parts := strings.Split(storedHash, "$")
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}
	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	// get input hash
	inputHash := argon2.IDKey([]byte(input), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	if subtle.ConstantTimeCompare(hash, inputHash) != 1 {
		return false, nil
	}
	return true, nil
}
