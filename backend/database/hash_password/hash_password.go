package hashpassword

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hashthispassword(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	bs := h.Sum(nil)

	return hex.EncodeToString(bs)
}
