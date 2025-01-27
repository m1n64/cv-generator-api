package utils

import (
	"crypto/sha256"
	"fmt"
)

func HashData(data []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(data))
}
