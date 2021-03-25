package key

import (
	"fmt"
)

func group(f string, v ...interface{}) string {
	return "[" + fmt.Sprintf(f, v...) + "]"
}
