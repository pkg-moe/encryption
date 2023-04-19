package encryption

import (
	"crypto/md5"
	"fmt"
)

func Md5(a string) string {
	has := md5.Sum([]byte(a))
	return fmt.Sprintf("%x", has)
}
