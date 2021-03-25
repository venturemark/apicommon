package key

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// User is the storage key for mapping user IDs. A user can have multiple
// associations with multiple resources.
func User(m map[string]string) *Key {
	var usi string
	{
		usi = m[metadata.UserID]
	}

	var k *Key
	{
		var id *ID
		if usi != "" {
			f, err := strconv.ParseFloat(usi, 64)
			if err != nil {
				panic(err)
			}
			s := usi

			id = &ID{
				f: f,
				s: s,
			}
		}

		var ele string
		{
			ele = fmt.Sprintf("use:%s", usi)
		}

		var lis string
		{
			lis = fmt.Sprintf("use")
		}

		k = &Key{
			id:  id,
			ele: ele,
			lis: lis,
		}
	}

	return k
}
