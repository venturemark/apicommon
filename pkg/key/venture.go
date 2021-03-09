package key

import (
	"fmt"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// Venture is the storage key for mapping venture IDs. A venture can have
// multiple timelines.
func Venture(m map[string]string) *Key {
	var ven string
	{
		ven = m[metadata.VentureID]
	}

	var k *Key
	{
		var ele string
		{
			ele = fmt.Sprintf("res:%s:ven", hash("ven:%s", ven))
		}

		var lis string
		{
			lis = fmt.Sprintf("res:%s:ven", hash("ven"))
		}

		var rol string
		{
			rol = fmt.Sprintf("res:%s:rol", hash(ele))
		}

		k = &Key{
			ele: ele,
			lis: lis,
			rol: rol,
		}
	}

	return k
}
