package utils

import (
	"crypto/sha256"
	"fmt"
)

func Hash256FromString(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)
	convertHash := fmt.Sprintf("%x", hash)
	return convertHash
}
