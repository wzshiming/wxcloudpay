package proto

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

func NonceStr() string {
	var buf [8]byte
	io.ReadFull(rand.Reader, buf[:])
	nonce := hex.EncodeToString(buf[:])
	return nonce
}
