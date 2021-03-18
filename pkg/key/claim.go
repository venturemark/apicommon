package key

import (
	"fmt"

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
				s: cli,
			}
		}

		var ele string
		{
			ele = fmt.Sprintf("res:%s:cla", hash("cla:%s", cli))
		}

		var lis string
		{
			lis = fmt.Sprintf("res:%s:cla", hash("cla"))
		}

		k = &Key{
			id:  id,
			ele: ele,
			lis: lis,
		}
	}

	return k
}
