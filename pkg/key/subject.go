package key

import (
	"fmt"

	"github.com/venturemark/apicommon/pkg/metadata"
)

// Subject is the storage key for mapping user IDs. A user can have multiple
// associations with multiple resources.
func Subject(m map[string]string) *Key {
	var rei string
	{
		switch m[metadata.ResourceKind] {
		case "invite":
			rei = Invite(m).List()
		case "message":
			rei = Message(m).List()
		case "timeline":
			rei = Timeline(m).List()
		case "update":
			rei = Update(m).List()
		case "user":
			rei = User(m).List()
		case "venture":
			rei = Venture(m).List()
		}
	}

	var sui string
	{
		sui = m[metadata.SubjectID]
	}

	var k *Key
	{
		var id *ID
		{
			id = &ID{
				s: sui,
			}
		}

		var ele string
		{
			ele = fmt.Sprintf("%s:sub:%s", rei, sui)
		}

		var lis string
		{
			lis = fmt.Sprintf("%s:sub", rei)
		}

		k = &Key{
			id:  id,
			ele: ele,
			lis: lis,
		}
	}

	return k
}
