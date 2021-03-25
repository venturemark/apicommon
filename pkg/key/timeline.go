package key

import (
	"fmt"
	"strconv"

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
		var id *ID
		if tii != "" {
			f, err := strconv.ParseFloat(tii, 64)
			if err != nil {
				panic(err)
			}
			s := tii

			id = &ID{
				f: f,
				s: s,
			}
		}

		var ele string
		{
			ele = fmt.Sprintf("ven:%s:tim:%s", vei, tii)
		}

		var lis string
		{
			lis = fmt.Sprintf("ven:%s:tim", vei)
		}

		k = &Key{
			id:  id,
			ele: ele,
			lis: lis,
		}
	}

	return k
}
