package encryption

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

func HmacSha1(input, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(input))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
