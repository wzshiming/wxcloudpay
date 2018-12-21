package authen

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func HmacSha256(data []byte, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	hash := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(hash))
}

func RSASSha256(data []byte, key crypto.Signer) (string, error) {
	data, err := key.Sign(rand.Reader, data, crypto.SHA256)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
