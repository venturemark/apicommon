package key

import (
	"strconv"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// Venture is the storage key for mapping venture IDs. A venture can have
// multiple timelines.
func Venture(m map[string]string) *Key {
	var vei string
	{
		vei = m[metadata.VentureID]
	}

	var k *Key
	{
		var id *ID
		if vei != "" {
			f, err := strconv.ParseFloat(vei, 64)
			if err != nil {
				panic(err)
			}
			s := vei

			id = &ID{
				f: f,
				s: s,
			}
		}

		var ele string
		{
			ele = group("res:%s:ven", group("ven:%s", vei))
		}

		var lis string
		{
			lis = group("res:%s:ven", group("ven"))
		}

		k = &Key{
			id:  id,
			ele: ele,
			lis: lis,
		}
	}

	return k
}
