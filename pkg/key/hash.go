package key

import (
	"crypto/sha256"
	"fmt"
)

func hash(f string, v ...interface{}) string {
	h := sha256.New()

	_, err := h.Write([]byte(fmt.Sprintf(f, v...)))
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
