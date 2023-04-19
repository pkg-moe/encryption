package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func Sha256(password string, salt string) string {
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(password))
	return fmt.Sprintf("%x", h.Sum(nil))
}
