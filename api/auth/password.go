package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Errors
var (
	errInvalidHash         = errors.New("the encoded hash is not in the correct format")
	errIncompatibleVersion = errors.New("incompatible version of argon2")
)

// passwordConfig configuración para generar el hash.
type passwordConfig struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint8
	keyLength   uint32
}

func generatePassword(c *passwordConfig, password string) (string, error) {
	// Generate a cryptographically secure random salt.
	salt, err := generateRandomBytes(c.saltLength)
	if err != nil {
		return "", err
	}

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash := argon2.IDKey([]byte(password), salt, c.iterations, c.memory, c.parallelism, c.keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b65Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodeHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, c.memory, c.iterations, c.parallelism, b64Salt, b65Hash,
	)
	return encodeHash, err
}

func comparePasswordAndHash(password, encodedHash string) (bool, error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	c, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey([]byte(password), salt, c.iterations, c.memory, c.parallelism, c.keyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	return subtle.ConstantTimeCompare(hash, otherHash) == 1, nil
}

func generateRandomBytes(n uint8) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func decodeHash(encodedHash string) (c *passwordConfig, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errInvalidHash
	}

	var version int
	if _, err = fmt.Sscanf(vals[2], "v=%d", &version); err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errIncompatibleVersion
	}

	c = &passwordConfig{}
	if _, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &c.memory, &c.iterations, &c.parallelism); err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	c.saltLength = uint8(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	c.keyLength = uint32(len(hash))

	return c, salt, hash, nil
}
