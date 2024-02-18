package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func HashRandomString(str string) (*string, error) {

	h := sha256.New()
	h.Write([]byte(str))
	random, err := GenerateRandomBytes(256)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	h.Write(random)

	sha := hex.EncodeToString(h.Sum(nil))

	return &sha, nil
}

func HashString(str string) (*string, error) {

	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	sha := hex.EncodeToString(bs)

	return &sha, nil
}
