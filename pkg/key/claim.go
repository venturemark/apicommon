package key

import (
	"github.com/venturemark/apicommon/pkg/metadata"
)

// Claim is the storage key for mapping user sibject claims to IDs. A user can
// have multiple subject claims.
func Claim(m map[string]string) *Key {
	var cli string
	{
		cli = m[metadata.ClaimID]
	}

	var k *Key
	{
		var id *ID
		{
			id = &ID{
				f: 0,
				s: cli,
			}
		}

		var ele string
		{
			ele = group("res:%s:cla", group("cla:%s", cli))
		}

		var lis string
		{
			lis = group("res:%s:cla", group("cla"))
		}

		k = &Key{
			id:  id,
			ele: ele,
			lis: lis,
		}
	}

	return k
}
