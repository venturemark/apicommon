package key

import (
	"fmt"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// Timeline is the storage key for mapping timeline IDs. A timeline can have
// multiple updates.
func Timeline(m map[string]string) *Key {
	var tii string
	{
		tii = m[metadata.TimelineID]
	}

	var vei string
	{
		vei = m[metadata.VentureID]
	}

	var k *Key
	{
		var ele string
		{
			ele = fmt.Sprintf("res:%s:tim", hash("ven:%s:tim:%s", vei, tii))
		}

		var lis string
		{
			lis = fmt.Sprintf("res:%s:tim", hash("ven:%s:tim", vei))
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
