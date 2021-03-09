package key

import (
	"fmt"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// Update is the storage key for mapping update IDs. An update can have multiple
// messages.
func Update(m map[string]string) *Key {
	var tii string
	{
		tii = m[metadata.TimelineID]
	}

	var upi string
	{
		upi = m[metadata.UpdateID]
	}

	var vei string
	{
		vei = m[metadata.VentureID]
	}

	var k *Key
	{
		var ele string
		{
			ele = fmt.Sprintf("res:%s:upd", hash("ven:%s:tim:%s:upd:%s", vei, tii, upi))
		}

		var lis string
		{
			lis = fmt.Sprintf("res:%s:upd", hash("ven:%s:tim:%s:upd", vei, tii))
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
