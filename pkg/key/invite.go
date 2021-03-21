package key

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// Invite is the storage key for mapping invite IDs. A user can be associated
// with one invite.
func Invite(m map[string]string) *Key {
	var ini string
	{
		ini = m[metadata.InviteID]
	}

	var vei string
	{
		vei = m[metadata.VentureID]
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
		{
			ele = fmt.Sprintf("res:%s:inv", hash("ven:%s:inv:%s", vei, ini))
		}

		var lis string
		{
			lis = fmt.Sprintf("res:%s:inv", hash("ven:%s:inv", vei))
		}

		k = &Key{
			id:  id,
			ele: ele,
			lis: lis,
		}
	}

	return k
}
