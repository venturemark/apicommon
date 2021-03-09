package key

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// Audience is the storage key for mapping audience IDs. A venture can have
// multiple audiences, while a user can be part of multiple ventures as well
// as multiple audiences.
func Audience(m map[string]string) *Key {
	var aui string
	{
		aui = m[metadata.AudienceID]
	}

	var k *Key
	{
		var id *ID
		{
			f, err := strconv.ParseFloat(aui, 64)
			if err != nil {
				panic(err)
			}
			s := aui

			id = &ID{
				f: f,
				s: s,
			}
		}

		var ele string
		{
			ele = fmt.Sprintf("res:%s:aud", hash("aud:%s", aui))
		}

		var lis string
		{
			lis = fmt.Sprintf("res:%s:aud", hash("aud"))
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
