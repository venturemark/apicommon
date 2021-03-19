package key

import (
	"fmt"
)

func hash(f string, v ...interface{}) string {
	return fmt.Sprintf(f, v...)

	// h := sha256.New()

	// _, err := h.Write([]byte(fmt.Sprintf(f, v...)))
	// if err != nil {
	// 	panic(err)
	// }

	// return fmt.Sprintf("%x", h.Sum(nil))
}
