package key

import (
	"fmt"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// Message is the storage key for mapping message IDs. A message can have
// multiple replies, which are just messages as well.
func Message(m map[string]string) *Key {
	var mei string
	{
		mei = m[metadata.MessageID]
	}

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
			ele = fmt.Sprintf("res:%s:mes", hash("ven:%s:tim:%s:upd:%s:mes:%s", vei, tii, upi, mei))
		}

		var lis string
		{
			lis = fmt.Sprintf("res:%s:mes", hash("ven:%s:tim:%s:upd:%s:mes", vei, tii, upi))
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
