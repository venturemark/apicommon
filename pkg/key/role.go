package key

import (
	"fmt"
	"strconv"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// Role is the storage key for mapping role IDs. A user can have multiple roles
// associated to resources. The inserted resource key is the filled, hashed
// version of any of the other resource keys defined here, e.g.
// sha265("ven:al9qy:tml") for records of timeline roles.
func Role(m map[string]string) *Key {
	var rei string
	{
		switch m[metadata.ResourceKind] {
		case "message":
			rei = Message(m).Elem()
		case "timeline":
			rei = Timeline(m).Elem()
		case "update":
			rei = Update(m).Elem()
		case "user":
			rei = User(m).Elem()
		case "venture":
			rei = Venture(m).Elem()
		}
	}

	var roi string
	{
		roi = m[metadata.RoleID]
	}

	var k *Key
	{
		var id *ID
		if roi != "" {
			f, err := strconv.ParseFloat(roi, 64)
			if err != nil {
				panic(err)
			}
			s := roi

			id = &ID{
				f: f,
				s: s,
			}
		}

		var ele string
		{
			ele = fmt.Sprintf("res:%s:rol", hash("res:%s:rol:%s", rei, roi))
		}

		var lis string
		{
			lis = fmt.Sprintf("res:%s:rol", hash("res:%s:rol", rei))
		}

		k = &Key{
			id:  id,
			ele: ele,
			lis: lis,
		}
	}

	return k
}
