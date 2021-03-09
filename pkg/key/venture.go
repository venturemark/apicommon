package key

import (
	"fmt"
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
		{
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
			ele = fmt.Sprintf("res:%s:ven", hash("ven:%s", vei))
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
			id:  id,
			ele: ele,
			lis: lis,
			rol: rol,
		}
	}

	return k
}
