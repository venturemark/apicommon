package key

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// Invite is the storage key for mapping invite IDs. A user can be associated
// with one invite per venture or timeline.
func Invite(m map[string]string) *Key {
	var ini string
	{
		ini = m[metadata.InviteID]
	}

	var vei string
	{
		vei = m[metadata.VentureID]
	}

	var tii string
	{
		tii = m[metadata.TimelineID]
	}

	var k *Key
	{
		var id *ID
		if ini != "" {
			f, err := strconv.ParseFloat(ini, 64)
			if err != nil {
				panic(err)
			}
			s := ini

			id = &ID{
				f: f,
				s: s,
			}
		}

		var ele string
		if tii == "" {
			ele = fmt.Sprintf("ven:%s:inv:%s", vei, ini)
		} else {
			ele = fmt.Sprintf("ven:%s:tim:%s:inv:%s", vei, tii, ini)
		}

		var lis string
		if tii == "" {
			lis = fmt.Sprintf("ven:%s:inv", vei)
		} else {
			lis = fmt.Sprintf("ven:%s:tim:%s:inv", vei, tii)
		}

		k = &Key{
			id:  id,
			ele: ele,
			lis: lis,
		}
	}

	return k
}
