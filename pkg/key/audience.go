package key

import (
	"fmt"

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
			ele: ele,
			lis: lis,
			rol: rol,
		}
	}

	return k
}
