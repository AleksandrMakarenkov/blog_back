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
		return false, err
	}
	return true, nil
}
