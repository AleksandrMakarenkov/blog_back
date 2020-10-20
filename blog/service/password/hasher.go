package password

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

const hashFormat = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"

func Hash(password string) (string, error) {
	p := &params{
		memory:      32 * 1024,
		iterations:  3,
		parallelism: 4,
		saltLength:  16,
		keyLength:   32,
	}

	salt := make([]byte, p.saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	saltBase64 := base64.RawStdEncoding.EncodeToString(salt)
	hashBase64 := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf(hashFormat, argon2.Version, p.memory, p.iterations, p.parallelism, saltBase64, hashBase64), nil
}
