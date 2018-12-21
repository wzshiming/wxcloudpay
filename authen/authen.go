package authen

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"strings"
)

func HmacSha256(data []byte, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	hash := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(hash))
}

func SignRSASSha256(data []byte, key crypto.Signer) (string, error) {
	data, err := key.Sign(rand.Reader, data, crypto.SHA256)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func GetPriKey(PriKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(PriKey)
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return priv, nil
	}
	priv2, err2 := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err2 != nil {
		return nil, err
	}
	priv3, ok := priv2.(*rsa.PrivateKey)
	if !ok {
		return nil, err
	}
	return priv3, nil
}
