package utils

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func EncryptPassword(ps string) (pwd string) {
	hash := sha3.Sum256([]byte(ps))
	pwd = hex.EncodeToString(hash[:])
	return pwd
}
