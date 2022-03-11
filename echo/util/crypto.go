package util

import (
	"crypto/sha1"
	"encoding/base64"
)

func DecodeSocketKey(secWebSocketKey string) string {
	var keyGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")
	h := sha1.New()
	h.Write([]byte(secWebSocketKey))
	h.Write(keyGUID)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
