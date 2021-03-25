package key

import (
	"fmt"
	"strconv"

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
		var id *ID
		if mei != "" {
			f, err := strconv.ParseFloat(mei, 64)
			if err != nil {
				panic(err)
			}
			s := mei

			id = &ID{
				f: f,
				s: s,
			}
		}

		var ele string
		{
			ele = fmt.Sprintf("ven:%s:tim:%s:upd:%s:mes:%s", vei, tii, upi, mei)
		}

		var lis string
		{
			lis = fmt.Sprintf("ven:%s:tim:%s:upd:%s:mes", vei, tii, upi)
		}

		k = &Key{
			id:  id,
			ele: ele,
			lis: lis,
		}
	}

	return k
}
