package encryption

import (
	"crypto/aes"
	"crypto/cipher"
)

// AESCtrCipher .
type AESCtrCipher struct {
	key   []byte
	nonce []byte
}

var iv = []byte("1s2z2SS0cg5a11a0")

// NewAESCtrCipher initializes AESCipher
func NewAESCtrCipher(key []byte) *AESCtrCipher {
	_, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	return &AESCtrCipher{key: key, nonce: iv}
}

// Encrypt 加密
func (a *AESCtrCipher) Encrypt(src []byte) []byte {
	if a == nil {
		return nil
	}
	cip, err := aes.NewCipher(a.key)
	if err != nil {
		return nil
	}
	blockMode := cipher.NewCTR(cip, a.nonce)
	res := make([]byte, len(src))
	blockMode.XORKeyStream(res, src)
	return res
}

// Decrypt 解密
func (a *AESCtrCipher) Decrypt(src []byte) []byte {
	if a == nil {
		return nil
	}
	cip, err := aes.NewCipher(a.key)
	if err != nil {
		return nil
	}
	blockMode := cipher.NewCTR(cip, a.nonce)
	res := make([]byte, len(src))
	blockMode.XORKeyStream(res, src)
	return res
}

// GetKey .
func (a *AESCtrCipher) GetKey() []byte {
	if a == nil {
		return nil
	}
	return a.key
}

// GetNonce .
func (a *AESCtrCipher) GetNonce() []byte {
	if a == nil {
		return nil
	}
	return a.nonce
}
